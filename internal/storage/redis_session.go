package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const sessionPrefix = "chat:session:"
const sessionTTL = 7 * 24 * time.Hour

type RedisSessionStorage struct {
	rdb *redis.Client
}

func NewRedisSessionStorage(addr string) *RedisSessionStorage {
	return &RedisSessionStorage{
		rdb: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

// CreateSession Создать сессию: token -> userID
func (s *RedisSessionStorage) CreateSession(ctx context.Context, token, userID string) error {
	return s.rdb.Set(ctx, sessionPrefix+token, userID, sessionTTL).Err()
}

// GetUserByToken Получить userID по токену
func (s *RedisSessionStorage) GetUserByToken(ctx context.Context, token string) (string, error) {
	return s.rdb.Get(ctx, sessionPrefix+token).Result()
}

// DeleteSession Удалить сессию (logout)
func (s *RedisSessionStorage) DeleteSession(ctx context.Context, token string) error {
	return s.rdb.Del(ctx, sessionPrefix+token).Err()
}
