package main

import (
	"bufio"
	"chat-grpc/api/pb"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	user := &pb.User{
		Id:       uuid.New().String(),
		Username: "test-user",
	}

	// Создаем поток
	stream, err := c.ChatStream(context.Background(), user)
	if err != nil {
		log.Fatalf("failed to call ChatStream: %v", err)
	}

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Println("stream recv err:", err)
				return
			}
			fmt.Printf("[%s] %s\n", msg.Username, msg.Text)
		}
	}()

	// Отправка сообщений с консоли
	fmt.Println("Connected. Type messages:")

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')

		_, err := c.SendMessage(context.Background(), &pb.MessageToServer{
			UserId:   user.Id,
			Username: user.Username,
			Text:     text,
		})
		if err != nil {
			log.Println("send error:", err)
		}
	}

}
