package handler

import (
	"net/http"

	"github.com/K31NER/notifier.git/internal/logger"
	"github.com/K31NER/notifier.git/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	url := h.authService.GetAuthURL()
	logger.Log.Info("Redirecting to Google Login")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) Callback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		logger.Log.Warn("OAuth callback received without code")
		c.JSON(http.StatusBadRequest, gin.H{"error": "code not found in query parameters"})
		return
	}

	logger.Log.Info("Exchanging authorization code for token")
	token, err := h.authService.ExchangeCode(c.Request.Context(), code)
	if err != nil {
		logger.Log.Error("Failed to exchange code for token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to exchange code for token", "details": err.Error()})
		return
	}

	logger.Log.Info("Successfully obtained OAuth tokens")
	c.JSON(http.StatusOK, gin.H{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"expiry":        token.Expiry,
		"token_type":    token.TokenType,
	})
}
