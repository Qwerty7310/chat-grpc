package storage

import (
	"chat-grpc/internal/models"
	"context"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}
