package provider

import (
	"context"
	"encoding/base64"

	"github.com/K31NER/notifier.git/internal/config"
	"github.com/K31NER/notifier.git/internal/logger"
	"github.com/K31NER/notifier.git/internal/models"
	"github.com/K31NER/notifier.git/internal/templates"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GmailProvider struct {
	service *gmail.Service
	sender  string
}

func NewGmailProvider(cfg *config.Config) *GmailProvider {
	ctx := context.Background()

	clientID := cfg.GMAIL_CLIENT_ID
	clientSecret := cfg.GMAIL_CLIENT_SECRET
	refreshToken := cfg.GMAIL_REFRESH_TOKEN
	sender := cfg.GMAIL_SENDER

	if clientID == "" || clientSecret == "" || refreshToken == "" {
		logger.Log.Warn("Gmail credentials are not fully set in config")
	}

	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{gmail.GmailSendScope},
	}

	// Al crear el token manualmente, es importante especificar el TokenType
	token := &oauth2.Token{
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	}

	// config.Client automáticamente se encarga de refrescar el access_token
	// cuando este expira, usando el refresh_token proporcionado.
	client := config.Client(ctx, token)

	service, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		logger.Log.Fatal("Failed to create gmail service", zap.Error(err))
	}

	logger.Log.Info("Gmail provider initialized successfully")

	return &GmailProvider{
		service: service,
		sender:  sender,
	}
}

func (g *GmailProvider) Send(ctx context.Context, email models.EmailRequest) error {
	logger.Log.Debug("Formatting email message", zap.String("to", email.Recipient))
	
	// Usar el paquete templates para construir el mensaje
	msgStr, err := templates.BuildEmailMessage(email)
	if err != nil {
		logger.Log.Error("Failed to build email message", zap.Error(err))
		return err
	}

	msg := &gmail.Message{
		Raw: base64.URLEncoding.EncodeToString([]byte(msgStr)),
	}

	logger.Log.Debug("Sending email via Gmail API")
	_, err = g.service.Users.Messages.Send("me", msg).Do()
	if err != nil {
		logger.Log.Error("Gmail API error during send", zap.Error(err))
		return err
	}
	
	return nil
}