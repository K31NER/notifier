package services

import (
	"context"

	"github.com/K31NER/notifier.git/internal/models"
	"github.com/K31NER/notifier.git/internal/provider"
)

type EmailService struct {
    provider provider.EmailProvider
}

func NewEmailService(p provider.EmailProvider) *EmailService {
    return &EmailService{provider: p}
}

func (s *EmailService) Send(ctx context.Context, email models.EmailRequest) error {
    return s.provider.Send(ctx, email)
}