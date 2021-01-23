package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/josemyduarte/grpc-go-demo/chat"
)

type Server struct {
	chat.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *chat.RequestMessage) (*chat.ResponseMessage, error) {
	log.Printf("Received message: [%s]", message)
	return &chat.ResponseMessage{Body: "This is the matrix"}, nil
}

func main() {
	const port = ":1111"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("connection to port %s: [%v]", port, err)
	}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Server failed to serve gRPC [%v]", err)
	}
}
