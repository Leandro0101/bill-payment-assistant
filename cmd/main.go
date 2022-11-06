package main

import (
	"bill-payment-assistant/server/ginserver"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	
	config := ginserver.ServerConfig{
		PORT: os.Getenv("APP_PORT"),
		APP_VERSION: os.Getenv("APP_VERSION"),
	}

	ginserver.Run(config)
}