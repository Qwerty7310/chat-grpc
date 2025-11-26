package storage

import (
	"chat-grpc/internal/models"
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	rdb *redis.Client
	key string
}

func NewRedisStorage(addr string, key string) *RedisStorage {
	r := redis.NewClient(&redis.Options{
		Addr: addr,
		//TODO: other options
	})

	return &RedisStorage{
		rdb: r,
		key: key,
	}
}

func (s *RedisStorage) AddMessage(ctx context.Context, msg *models.Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if err := s.rdb.LPush(ctx, s.key, data).Err(); err != nil {
		return err
	}

	if err := s.rdb.LTrim(ctx, s.key, 0, 99).Err(); err != nil {
		return err
	}

	return nil
}

func (s *RedisStorage) GetLastMessages(ctx context.Context, limit int) ([]*models.Message, error) {
	values, err := s.rdb.LRange(ctx, s.key, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	msgs := make([]*models.Message, 0, len(values))
	for _, v := range values {
		var m models.Message
		if err := json.Unmarshal([]byte(v), &m); err != nil {
			continue
		}
		msgs = append(msgs, &m)
	}
	return msgs, nil
}

func (s *RedisStorage) Ping(ctx context.Context) error {
	return s.rdb.Ping(ctx).Err()
}
