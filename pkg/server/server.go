package server

import (
	"context"
	"fmt"
	"github.com/gorobot-nz/go-grpc-task/pkg/bitlyClient"
	challenge "github.com/gorobot-nz/go-grpc-task/pkg/gen/pkg/proto"
	"google.golang.org/grpc"
	"net"
)

type challengeServiceServer struct {
	challenge.UnimplementedChallengeServiceServer
	client *bitlyClient.BitlyClient
}

func NewServer(bClient *bitlyClient.BitlyClient) *grpc.Server {
	server := grpc.NewServer()
	challenge.RegisterChallengeServiceServer(server, &challengeServiceServer{client: bClient})
	return server
}

func Run(server *grpc.Server, port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", port, err)
	}
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}
	return nil
}

func (s *challengeServiceServer) MakeShortLink(ctx context.Context, link *challenge.Link) (*challenge.Link, error) {
	data := link.GetData()
	resp, err := s.client.ShortLink(data)
	if err != nil {
		return nil, err
	}

	return &challenge.Link{Data: resp}, nil
}
