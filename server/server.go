package main

import (
	Chat "Golang_Chat_System/Chat"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type ChatServer struct {
	Chat.UnimplementedChattingServiceServer
}

var users = make(map[string]Chat.ChattingService_JoinChatServer)

func main() {
	listen, err := net.Listen("tcp", ":8007")
	if err != nil {
		log.Fatalf("Failed to listen on port 8007: %v", err)
	}
	log.Println(":8007 is listening")

	// Creates empty gRPC server
	grpcServer := grpc.NewServer()

	// Creates instance of our ChatServiceServer struct and binds it with our empty gRPC server.
	ccs := ChatServer{}
	Chat.RegisterChattingServiceServer(grpcServer, &ccs)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC server :: %v", err)
	}
}

func (c *ChatServer) JoinChat(user *Chat.User, ccsi Chat.ChattingService_JoinChatServer) error {

	users[user.Name] = ccsi

	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic: %v", err)
			os.Exit(1)
		}
	}()

	body := fmt.Sprintf(user.Name + " has joined the chat. The new user's public key should now be validated by the server via. Public Key Autherization.")

	Broadcast("ServerMessage", body, []byte{})

	// block function
	bl := make(chan bool)
	<-bl

	return nil
}

func (is *ChatServer) SendEncrypted(ctx context.Context, msg *Chat.ClientEncrypted) (*Chat.Empty, error) {
	Broadcast(msg.Name, msg.Message, msg.Signature)
	return &Chat.Empty{}, nil
}

func Broadcast(name string, body string, signature []byte) {

	log.Printf("%s : %s)", name, body)

	for key, value := range users {
		if key != name {
			err := value.Send(&Chat.FromServer{Name: name, Body: body, Signature: signature})
			if err != nil {
				log.Println("Failed to broadcast to "+string(key)+": ", err)
			}
		}
	}
}
