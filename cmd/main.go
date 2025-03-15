package main

import (
	"fmt"
	"os"

	"github.com/anashasan312/movie-reservation-system/cmd/server"
	log "github.com/anashasan312/movie-reservation-system/logger"
	"github.com/anashasan312/movie-reservation-system/pkg/infrastructure/config"
)

func main() {

	fmt.Println("Starting...")

	var (
		env            = os.Getenv("ENVIRONMENT")
		configFilePath = os.Getenv("CONFIG_FILE_PATH")
		secretFilePath = os.Getenv("SECRET_FILE_PATH")
	)

	log.Boot(fmt.Sprintf("ENVIRONMENT:%s, CONFIG_FILE_PATH: %s, SECRET_FILE_PATH: %s", env, configFilePath, secretFilePath))

	conf, err := config.LoadConfig(configFilePath, secretFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read config/secret file, err: %v. Shutting down.", err.Error()))
	}

	s, err := server.NewServer(conf)

	if err != nil {
		panic(fmt.Sprintf("Failed to start server, err:%v. Shutting down.", err.Error()))
	}

	s.SetupRoutes()

	s.Run()
}
