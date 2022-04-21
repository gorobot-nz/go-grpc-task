package main

import (
	"github.com/gorobot-nz/go-grpc-task/pkg/client"
	"log"
)

func main() {
	challengeClient := client.NewClient("127.0.0.1:8080")
	err := challengeClient.Run()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
