package server

import (
	"chat-grpc/api/pb"
	"context"

	"sync"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type clientStream struct {
	userId string
	stream pb.ChatService_ChatStreamServer
}

type ChatService struct {
	pb.UnimplementedChatServiceServer

	mu       sync.RWMutex
	clients  map[string]clientStream // userID -> stream
	messages []*pb.Message           // in-memory history
}

func NewChatService() *ChatService {
	return &ChatService{
		clients:  make(map[string]clientStream),
		messages: make([]*pb.Message, 0),
	}
}

func (s *ChatService) ChatStream(user *pb.User, stream pb.ChatService_ChatStreamServer) error {
	s.mu.Lock()
	s.clients[user.Id] = clientStream{
		userId: user.Id,
		stream: stream,
	}
	s.mu.Unlock()

	// history
	// ...

	<-stream.Context().Done()

	s.mu.Lock()
	delete(s.clients, user.Id)
	s.mu.Unlock()

	return nil
}

func (s *ChatService) broadcast(msg *pb.Message) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, client := range s.clients {
		go func(c clientStream) {
			if err := c.stream.Send(msg); err != nil {
				s.mu.Lock()
				delete(s.clients, c.userId)
				s.mu.Unlock()
			}
		}(client)
	}
}

func (s *ChatService) SendMessage(ctx context.Context, req *pb.MessageToServer) (*pb.SendAck, error) {
	id := uuid.New().String()
	now := timestamppb.Now()

	msg := &pb.Message{
		Id:        id,
		UserId:    req.UserId,
		Username:  req.Username,
		Text:      req.Text,
		CreatedAt: now,
	}

	s.mu.Lock()
	s.messages = append(s.messages, msg)
	s.mu.Unlock()

	s.broadcast(msg)

	return &pb.SendAck{
		Ok: true,
		Id: id,
	}, nil
}
