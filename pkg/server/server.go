package server

import (
	"context"
	"fmt"
	"github.com/gorobot-nz/go-grpc-task/pkg/apiClients"
	challenge "github.com/gorobot-nz/go-grpc-task/pkg/gen/pkg/proto"
	"github.com/zpnk/go-bitly"
	"google.golang.org/grpc"
	"net"
	"time"
)

type timerSubscribers struct {
	timer       *challenge.Timer
	subscribers []*challenge.ChallengeService_StartTimerServer
}

type challengeServiceServer struct {
	challenge.UnimplementedChallengeServiceServer
	bClient     *bitly.Client
	timerClient *apiclients.TimerClient
}

func NewServer(bClient *bitly.Client, tClient *apiclients.TimerClient) *grpc.Server {
	server := grpc.NewServer()
	challenge.RegisterChallengeServiceServer(server, &challengeServiceServer{bClient: bClient, timerClient: tClient})
	return server
}

func Run(server *grpc.Server, port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", port, err)
	}
	fmt.Println("Server start")
	fmt.Sprintf("Port start %s", port)
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
	resp, err := s.bClient.Links.Shorten(data)
	if err != nil {
		return nil, err
	}

	return &challenge.Link{Data: resp.URL}, nil
}

func (s *challengeServiceServer) StartTimer(timer *challenge.Timer, timerServer challenge.ChallengeService_StartTimerServer) error {
	err := s.timerClient.CreateTimer(timer)
	if err != nil {
		return err
	}
	for {
		seconds, err := s.timerClient.GetRemainingSeconds(timer)
		fmt.Println("Get timer")
		if err != nil {
			break
		}
		f, _ := seconds.Int64()
		err = timerServer.Send(&challenge.Timer{Name: timer.Name, Seconds: f})
		time.Sleep(time.Duration(timer.Frequency) * time.Second)
	}
	return nil
}
