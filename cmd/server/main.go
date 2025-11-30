package main

import (
	"chat-grpc/api/pb"
	"chat-grpc/internal/server"
	"chat-grpc/internal/storage"
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	//if err := godotenv.Load(); err != nil {
	//	log.Println("No .env file found, using environment variables")
	//}

	pg, err := storage.NewPostgresStorage(os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	log.Println("Postgres connected.")

	redisCache := storage.NewRedisStorage(os.Getenv("REDIS_ADDR"), "chat:messages")

	sessionStore := storage.NewRedisSessionStorage(os.Getenv("REDIS_ADDR"))

	if err := redisCache.Ping(context.Background()); err != nil {
		log.Fatalf("redis unavailable: %v", err)
	}
	log.Println("Redis connected.")

	store := storage.NewHybridStorage(redisCache, pg)

	chatService := server.NewChatService(store)

	authService := server.NewAuthService(pg, sessionStore)

	interceptor := server.NewAuthInterceptor(sessionStore)

	_ = sessionStore

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	pb.RegisterChatServiceServer(s, chatService)
	pb.RegisterAuthServiceServer(s, authService)

	log.Println("gRPC server listening on :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
