package storage

import (
	"chat-grpc/internal/models"
)

type Storage interface {
	AddMessage(msg *models.Message) error
	GetLastMessages(limit int) ([]*models.Message, error)
}
