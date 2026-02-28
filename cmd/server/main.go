package main

import (
	"fmt"

	"github.com/K31NER/notifier.git/internal/config"
	"github.com/K31NER/notifier.git/internal/handler"
	"github.com/K31NER/notifier.git/internal/logger"
	"github.com/K31NER/notifier.git/internal/provider"
	service "github.com/K31NER/notifier.git/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Inicializacion del logger
	logger.InitLogger()
	defer logger.Sync() 

	// Cargamos las varibales de entorno
	cfg := config.Load()

	logger.Log.Info("Starting Email Sender Service...")

	// Obtenemos el modo del servidor
	if cfg.APP_ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	gmailProvider := provider.NewGmailProvider(cfg)
	emailService := service.NewEmailService(gmailProvider)
	emailHandler := handler.NewEmailHandler(emailService)

	// Rutas
	r.POST("/send", emailHandler.Send)

	logger.Log.Info(fmt.Sprintf("Server listening on port %s", cfg.PORT))
	if err := r.Run(":" + cfg.PORT); err != nil {
		logger.Log.Fatal("Failed to start server", zap.Error(err))
	}
}