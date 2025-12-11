package main

import (
	"bufio"
	"chat-grpc/api/pb"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	authClient := pb.NewAuthServiceClient(conn)
	chatClient := pb.NewChatServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)

	// ---------------------------
	//   1. LOGIN / REGISTER
	// ---------------------------

	fmt.Println("Login on Register (l/r)?")
	ch, _ := reader.ReadString('\n')
	ch = strings.TrimSpace(ch)

	var token string
	var userID string

	if ch == "r" {
		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		resp, err := authClient.Register(context.Background(), &pb.RegisterRequest{
			Username: username,
			Password: password,
		})

		if err != nil {
			log.Fatalf("Failed to register user: %v", err)
		}

		if !resp.Ok {
			log.Fatalf("Failed to register user: %v", resp.Error)
		}

		token = resp.Token
		userID = resp.UserId
	} else {
		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		resp, err := authClient.Login(context.Background(), &pb.LoginRequest{
			Username: username,
			Password: password,
		})

		if err != nil {
			log.Fatalf("Failed to login: %v", err)
		}

		if !resp.Ok {
			log.Fatalf("Failed to login: %v", resp.Error)
		}

		token = resp.Token
		userID = resp.UserId
	}

	fmt.Println("Logged in as:", userID)

	// ---------------------------
	//       2. AUTH CONTEXT
	// ---------------------------

	authCtx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("authorization", "Bearer "+token))

	// ---------------------------
	//       3. START STREAM
	// ---------------------------

	stream, err := chatClient.ChatStream(authCtx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Failed to ChatStream: %v", err)
	}

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Println("stream recv error:", err)
				return
			}

			fmt.Printf("[%s] %s\n", msg.Username, msg.Text)
		}
	}()

	// ---------------------------
	//       4. SEND MESSAGES
	// ---------------------------

	fmt.Println("Connected. Type messages:")

	for {
		text, _ := reader.ReadString('\n')
		text = sanitizeUTF8(text)

		_, err := chatClient.SendMessage(authCtx, &pb.MessageToServer{
			Text: strings.TrimSpace(text),
		})

		if err != nil {
			log.Println("send error:", err)
		}
	}
}

func sanitizeUTF8(s string) string {
	var b []rune
	for _, r := range s {
		if r == '\uFFFD' { // invalid rune
			continue
		}
		b = append(b, r)
	}
	return string(b)
}
