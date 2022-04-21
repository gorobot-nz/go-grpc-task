package main

import (
	"github.com/gorobot-nz/go-grpc-task/pkg/client"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	challengeClient := client.NewClient(viper.GetString("HOST"))
	err := challengeClient.Run()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
