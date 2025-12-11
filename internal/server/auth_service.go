package server

import (
	"chat-grpc/api/pb"
	"chat-grpc/internal/models"
	"chat-grpc/internal/storage"
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	users   storage.UserStorage
	session *storage.RedisSessionStorage
}

func NewAuthService(users storage.UserStorage, session *storage.RedisSessionStorage) *AuthService {
	return &AuthService{
		users:   users,
		session: session,
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()

	// Проверка на существование такого пользователя
	existing, err := s.users.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return &pb.AuthResponse{
			Ok:    false,
			Error: "username already exists",
		}, nil
	}

	// хэшируем пароль
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// создание пользователя
	user := &models.User{
		ID:           uuid.New().String(),
		Username:     username,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
	}

	if err := s.users.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	// создаем токен
	token := uuid.New().String()

	if err := s.session.CreateSession(ctx, token, user.ID); err != nil {
		return nil, err
	}

	return &pb.AuthResponse{
		Ok:     true,
		Token:  token,
		UserId: user.ID,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()

	user, err := s.users.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return &pb.AuthResponse{
			Ok:    false,
			Error: "invalid username or password",
		}, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return &pb.AuthResponse{
			Ok:    false,
			Error: "invalid username or password",
		}, nil
	}

	token := uuid.New().String()
	if err := s.session.CreateSession(ctx, token, user.ID); err != nil {
		return nil, err
	}

	return &pb.AuthResponse{
		Ok:     true,
		Token:  token,
		UserId: user.ID,
	}, nil
}
