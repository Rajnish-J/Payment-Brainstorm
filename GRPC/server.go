package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    "GRPC/GRPC"
)

func main() {
    lis, err := net.Listen("tcp", ":9000")
    if err != nil {
        log.Fatalf("Failed to listen on port :9000 %v", err)
    }

    s := GRPC.Server{}

    grpcServer := grpc.NewServer()

    GRPC.RegisterChatServiceServer(grpcServer, &s)

    log.Println("gRPC server started on :9000")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve GRPC server on port :9000 %v", err)
    }
}
