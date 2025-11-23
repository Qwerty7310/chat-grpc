package models

import (
	"chat-grpc/api/pb"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Message struct {
	ID       string
	UserID   string
	Username string
	Text     string
	CreateAt time.Time
}

func FromProto(p *pb.Message) *Message {
	return &Message{
		ID:       p.Id,
		UserID:   p.UserId,
		Username: p.Username,
		Text:     p.Text,
		CreateAt: p.CreatedAt.AsTime(),
	}
}

func (m *Message) ToProto() *pb.Message {
	return &pb.Message{
		Id:        m.ID,
		UserId:    m.UserID,
		Username:  m.Username,
		Text:      m.Text,
		CreatedAt: timestamppb.New(m.CreateAt),
	}
}
