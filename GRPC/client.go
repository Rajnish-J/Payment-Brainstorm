package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"GRPC/GRPC"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	c := GRPC.NewChatServiceClient(conn)

	message := GRPC.Message{
		Body: "Hello from the client",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling sayHello: %s", err)
	}

	log.Printf("Response from the server: %s", response.Body)
}