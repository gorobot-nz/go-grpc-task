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
	bClient *bitlyClient.BitlyClient
}

func NewServer(bClient *bitlyClient.BitlyClient) *grpc.Server {
	server := grpc.NewServer()
	challenge.RegisterChallengeServiceServer(server, &challengeServiceServer{bClient: bClient})
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

func (s *challengeServiceServer) ReadMetadata(ctx context.Context, placeholder *challenge.Placeholder) (*challenge.Placeholder, error) {
	metadata := placeholder.GetData()
	return &challenge.Placeholder{Data: metadata}, nil
}

func (s *challengeServiceServer) MakeShortLink(ctx context.Context, link *challenge.Link) (*challenge.Link, error) {
	data := link.GetData()
	resp, err := s.bClient.ShortLink(data)
	if err != nil {
		return nil, err
	}

	return &challenge.Link{Data: resp}, nil
}
