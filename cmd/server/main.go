package main

import (
	"github.com/gorobot-nz/go-grpc-task/pkg/apiClients"
	serverPkg "github.com/gorobot-nz/go-grpc-task/pkg/server"
	"github.com/spf13/viper"
	"github.com/zpnk/go-bitly"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	bitlyClient := bitly.New(viper.GetString("BITLY_OAUTH_TOKEN"))
	timerClient := apiclients.NewTimerClient()
	server := serverPkg.NewServer(bitlyClient, timerClient)
	err := serverPkg.Run(server, viper.GetString("HOST"))
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
