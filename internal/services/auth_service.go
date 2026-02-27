package services

import (
	"context"

	"github.com/K31NER/notifier.git/internal/config"
	"github.com/K31NER/notifier.git/internal/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

type AuthService struct {
	oauthConfig *oauth2.Config
}

func NewAuthService(cfg *config.Config) *AuthService {
	clientID := cfg.GMAIL_CLIENT_ID
	clientSecret := cfg.GMAIL_CLIENT_SECRET

	if clientID == "" || clientSecret == "" {
		logger.Log.Warn("GMAIL_CLIENT_ID or GMAIL_CLIENT_SECRET is not set")
	}

	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/oauth2callback",
		Scopes:       []string{gmail.GmailSendScope},
	}

	logger.Log.Info("AuthService initialized")

	return &AuthService{
		oauthConfig: config,
	}
}

func (s *AuthService) GetAuthURL() string {
	return s.oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("prompt", "consent"))
}

func (s *AuthService) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	return s.oauthConfig.Exchange(ctx, code)
}
