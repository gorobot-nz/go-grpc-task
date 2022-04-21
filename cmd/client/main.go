package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"

	challenges "github.com/gorobot-nz/go-grpc-task/pkg/gen/pkg/proto"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
func run() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to PetStoreService on %s: %w", connectTo, err)
	}
	log.Println("Connected to", connectTo)
	challengeClient := challenges.NewChallengeServiceClient(conn)
	stream, err := challengeClient.StartTimer(context.Background(), &challenges.Timer{Name: "test", Seconds: 60, Frequency: 5})
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", challengeClient, err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}

		fmt.Println(res.GetSeconds())
	}
	return nil
}
