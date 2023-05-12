package main

import (
	"log"

	"github.com/BrazenFox/compiler-service/internal/app/handler"
	"github.com/BrazenFox/compiler-service/internal/app/server"
	"github.com/BrazenFox/compiler-service/internal/app/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("err: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
