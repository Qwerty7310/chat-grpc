package storage

import (
	"chat-grpc/internal/models"
	"database/sql"
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

func (s *PostgresStorage) AddMessage(msg *models.Message) error {
	_, err := s.db.Exec(`
		INSERT INTO messages (id, user_id, username, text, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`, msg.ID, msg.UserID, msg.Username, msg.Text, msg.CreateAt)
	return err
}

func (s *PostgresStorage) GetLastMessages(limit int) ([]*models.Message, error) {
	rows, err := s.db.Query(`
		SELECT id, user_id, username, text, created_at
		FROM messages
		ORDER BY created_at DESC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []*models.Message
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.ID, &m.UserID, &m.Username, &m.Text, &m.CreateAt); err != nil {
			return nil, err
		}
		msgs = append(msgs, &m)
	}
	return msgs, nil
}
