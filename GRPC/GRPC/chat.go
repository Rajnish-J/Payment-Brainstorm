package GRPC

import (
	"log"
	"golang.org/x/net/context"
)

type Server struct{
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Receive Message  body from client: %s", message.Body)
	return &Message{Body: "Hello from the server"}, nil
}