package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	chat "github.com/arivlav/chat-server/pkg/chat_v1"
)

const (
	grpcPort = 50052
)

type server struct {
	chat.UnimplementedChatV1Server
}

// Create ...
func (s *server) Create(_ context.Context, req *chat.CreateRequest) (*chat.CreateResponse, error) {
	id := gofakeit.Uint64()
	log.Printf("New chat (%d) is created: ", id)
	log.Printf("Users: %+v", req.GetUsernames())

	return &chat.CreateResponse{
		Id: id,
	}, nil
}

// Delete ...
func (s *server) Delete(_ context.Context, req *chat.DeleteRequest) (*empty.Empty, error) {
	log.Printf("Chat %d is deleted", req.GetId())

	return &empty.Empty{}, nil
}

// SendMessage ...
func (s *server) SendMessage(_ context.Context, req *chat.SendMessageRequest) (*empty.Empty, error) {
	log.Printf("Send message:")
	log.Printf("from: %v", req.GetFrom())
	log.Printf("text: %v", req.GetText())

	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
