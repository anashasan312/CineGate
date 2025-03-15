package main

import (
	"fmt"
	"os"

	"github.com/InstaViewAI/genie/logger/log"
)

func main() {

	fmt.Println("Starting...")

	var (
		env            = os.Getenv("ENVIRONMENT")
		configFilePath = os.Getenv("CONFIG_FILE_PATH")
		secretFilePath = os.Getenv("SECRET_FILE_PATH")
	)

	log.Boot(fmt.Sprintf("ENVIRONMENT:%s, CONFIG_FILE_PATH: %s, SECRET_FILE_PATH: %s", env, configFilePath, secretFilePath))
}
