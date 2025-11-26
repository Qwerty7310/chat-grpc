package main

import (
	"chat-grpc/api/pb"
	"chat-grpc/internal/server"
	"chat-grpc/internal/storage"
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	pg, err := storage.NewPostgresStorage(os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	log.Println("Postgres connected.")

	redis := storage.NewRedisStorage(os.Getenv("REDIS_ADDR"), "chat:messages")

	if err := redis.Ping(context.Background()); err != nil {
		log.Fatalf("redis unavailable: %v", err)
	}
	log.Println("Redis connected.")

	store := storage.NewHybridStorage(redis, pg)

	chatService := server.NewChatService(store)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, chatService)

	log.Println("gRPC server listening on :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
