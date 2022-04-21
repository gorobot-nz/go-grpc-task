package main

import (
	"github.com/gorobot-nz/go-grpc-task/pkg/apiClients"
	serverPkg "github.com/gorobot-nz/go-grpc-task/pkg/server"
	"github.com/zpnk/go-bitly"
	"log"
)

func main() {
	bitlyClient := bitly.New("2f6a141a4b9961f2424e76d35b19b126b1368852")
	timerClient := apiclients.NewTimerClient()
	server := serverPkg.NewServer(bitlyClient, timerClient)
	err := serverPkg.Run(server, "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}
