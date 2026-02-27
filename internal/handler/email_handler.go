package handler

import (
	"context"
	"net/http"

	"github.com/K31NER/notifier.git/internal/logger"
	"github.com/K31NER/notifier.git/internal/models"
	service "github.com/K31NER/notifier.git/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EmailHandler struct {
	service *service.EmailService
}

func NewEmailHandler(s *service.EmailService) *EmailHandler {
	return &EmailHandler{service: s}
}

func (h *EmailHandler) Send(c *gin.Context) {
	var req models.EmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Warn("Invalid request payload", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Log.Info("Queuing email for sending", zap.String("recipient", req.Recipient), zap.String("subject", req.Subject))

	// Lanzar una goroutine para enviar el correo de forma asíncrona
	go func(request models.EmailRequest) {
		// Creamos un contexto independiente porque el del request se cancela al terminar la respuesta HTTP
		ctx := context.Background()
		if err := h.service.Send(ctx, request); err != nil {
			logger.Log.Error("Failed to send email asynchronously", zap.Error(err), zap.String("recipient", request.Recipient))
			return
		}
		logger.Log.Info("Email sent successfully", zap.String("recipient", request.Recipient))
	}(req)

	// Responder inmediatamente al cliente
	c.JSON(http.StatusAccepted, gin.H{"status": "email queued", "message": "The email is being processed in the background"})
}