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

var users = make(map[int32]Chat.ChattingService_JoinChatServer)

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

	users[user.Id] = ccsi

	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic: %v", err)
			os.Exit(1)
		}
	}()

	body := fmt.Sprintf(user.Name + " has joined the chat. The new user's public key should now be validated by the server via. Public Key Autherization.")

	Broadcast(&Chat.ClientContent{Name: "ServerMessage", Body: body})

	// block function
	bl := make(chan bool)
	<-bl

	return nil
}

func (is *ChatServer) SendContent(ctx context.Context, msg *Chat.ClientContent) (*Chat.Empty, error) {

	Broadcast(msg)
	return &Chat.Empty{}, nil
}

func Broadcast(msg *Chat.ClientContent) {
	name := msg.Name
	body := msg.Body

	log.Printf("%s : %s", name, body)

	for key, value := range users {
		err := value.Send(&Chat.FromServer{Name: name, Body: body})
		if err != nil {
			log.Println("Failed to broadcast to "+string(key)+": ", err)
		}
	}
}

func (is *ChatServer) RevealAll(ctx context.Context, msg *Chat.ClientRevelation) (*Chat.Empty, error) {

	InfoBroadcast(msg)
	return &Chat.Empty{}, nil
}

func InfoBroadcast(msg *Chat.ClientRevelation) {
	name := msg.Name
	c := msg.C
	m := msg.M
	r := msg.R
	result := msg.Result

	log.Printf("%s : [%d, %d, %d, %d]", name, c, m, r, result)

	for key, value := range users {
		err := value.Send(&Chat.FromServer{Name: name, Body: fmt.Sprintf("[%d, %d, %d, %d]", c, m, r, result)})
		if err != nil {
			log.Println("Failed to broadcast to "+string(key)+": ", err)
		}
	}
}
