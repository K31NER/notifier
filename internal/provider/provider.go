package provider

import (
	"context"

	"github.com/K31NER/notifier.git/internal/models"
)

type EmailProvider interface {
    Send(ctx context.Context, email models.EmailRequest) error
}