package storage

import (
	"chat-grpc/internal/models"
	"context"
)

type Storage interface {
	AddMessage(ctx context.Context, msg *models.Message) error
	GetLastMessages(ctx context.Context, limit int) ([]*models.Message, error)
}
