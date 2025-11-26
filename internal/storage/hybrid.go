package storage

import (
	"chat-grpc/internal/models"
	"context"
	"log"
)

type HybridStorage struct {
	redis *RedisStorage
	pg    *PostgresStorage
}

func NewHybridStorage(redis *RedisStorage, pg *PostgresStorage) *HybridStorage {
	return &HybridStorage{redis: redis, pg: pg}
}

func (s *HybridStorage) AddMessage(ctx context.Context, msg *models.Message) error {
	if err := s.pg.AddMessage(ctx, msg); err != nil {
		return err
	}

	if err := s.redis.AddMessage(ctx, msg); err != nil {
		log.Printf("redis cache update failed: %v", err)
	}

	return nil
}

func (s *HybridStorage) GetLastMessages(ctx context.Context, limit int) ([]*models.Message, error) {
	msgs, err := s.redis.GetLastMessages(ctx, limit)
	if err == nil && len(msgs) > 0 {
		return msgs, nil
	}

	dbMsgs, err := s.pg.GetLastMessages(ctx, limit)
	if err != nil {
		return nil, err
	}

	for _, m := range dbMsgs {
		s.redis.AddMessage(ctx, m)
	}

	return dbMsgs, nil
}
