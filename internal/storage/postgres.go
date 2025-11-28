package storage

import (
	"chat-grpc/internal/models"
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) DB() *sql.DB {
	return s.db
}

func (s *PostgresStorage) AddMessage(ctx context.Context, msg *models.Message) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO messages (id, user_id, username, text, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`, msg.ID, msg.UserID, msg.Username, msg.Text, msg.CreatedAt)
	return err
}

func (s *PostgresStorage) GetLastMessages(ctx context.Context, limit int) ([]*models.Message, error) {
	rows, err := s.db.QueryContext(ctx, `
		SELECT id, user_id, username, text, created_at
		FROM messages
		ORDER BY created_at DESC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	msgs := make([]*models.Message, 0, limit)
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.ID, &m.UserID, &m.Username, &m.Text, &m.CreatedAt); err != nil {
			return nil, err
		}
		msgs = append(msgs, &m)
	}
	return msgs, nil
}

func (s *PostgresStorage) CreateUser(ctx context.Context, user *models.User) error {
	_, err := s.db.ExecContext(ctx,
		`INSERT INTO users (id, username, password_hash, created_at)
			   VALUES ($1, $2, $3, $4)`,
		user.ID, user.Username, user.PasswordHash, user.CreatedAt,
	)
	return err
}

func (s *PostgresStorage) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	row := s.db.QueryRowContext(ctx,
		`SELECT id, username, password_hash, created_at FROM users WHERE username = $1`, username)

	var u models.User
	if err := row.Scan(&u.ID, &u.Username, &u.PasswordHash, &u.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}
