package main

import (
	"github.com/gorobot-nz/go-grpc-task/pkg/bitlyClient"
	serverPkg "github.com/gorobot-nz/go-grpc-task/pkg/server"
	"log"
)

func main() {
	bitlyClient := bitlyClient.NewBitlyClient("49531020c64e26d600955a9c6d4f198a3543d7c0")
	server := serverPkg.NewServer(bitlyClient)
	err := serverPkg.Run(server, "8080")
	if err != nil {
		log.Fatal(err)
	}
}
