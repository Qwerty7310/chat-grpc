package server

import (
	"chat-grpc/api/pb"
	"chat-grpc/internal/models"
	"chat-grpc/internal/storage"
	"context"

	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type clientStream struct {
	userId string
	stream pb.ChatService_ChatStreamServer
}

type ChatService struct {
	pb.UnimplementedChatServiceServer

	mu      sync.RWMutex
	clients map[string]clientStream // userID -> stream
	store   storage.Storage
}

func NewChatService(store storage.Storage) *ChatService {
	return &ChatService{
		clients: make(map[string]clientStream),
		store:   store,
	}
}

func (s *ChatService) ChatStream(_ *emptypb.Empty, stream pb.ChatService_ChatStreamServer) error {
	ctx := stream.Context()

	uidValue := ctx.Value(ctxKeyUserID)
	if uidValue == nil {
		return status.Error(codes.PermissionDenied, "user not authorized")
	}
	uid := uidValue.(string)

	user, err := s.store.GetUserByID(ctx, uid)
	if err != nil {
		return status.Error(codes.Internal, "failed to load user")
	}

	_ = user

	s.mu.Lock()
	s.clients[uid] = clientStream{
		userId: uid,
		stream: stream,
	}
	s.mu.Unlock()

	// history
	// ...

	<-stream.Context().Done()

	s.mu.Lock()
	delete(s.clients, uid)
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
	uidValue := ctx.Value(ctxKeyUserID)
	if uidValue == nil {
		return nil, status.Error(codes.PermissionDenied, "user not authorized")
	}
	uid := uidValue.(string)

	user, err := s.store.GetUserByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	msg := &models.Message{
		ID:        uuid.New().String(),
		UserID:    uid,
		Username:  user.Username,
		Text:      req.Text,
		CreatedAt: timestamppb.Now().AsTime(),
	}

	if err := s.store.AddMessage(ctx, msg); err != nil {
		return nil, err
	}

	pbMsg := msg.ToProto()
	s.broadcast(pbMsg)

	return &pb.SendAck{
		Ok: true,
		Id: msg.ID,
	}, nil
}

func (s *ChatService) GetHistory(ctx context.Context, req *pb.HistoryRequest) (*pb.GetHistoryResponse, error) {
	msgs, err := s.store.GetLastMessages(ctx, int(req.Limit))
	if err != nil {
		return nil, err
	}

	res := make([]*pb.Message, 0, len(msgs))
	for _, m := range msgs {
		res = append(res, m.ToProto())
	}

	return &pb.GetHistoryResponse{Message: res}, nil
}
