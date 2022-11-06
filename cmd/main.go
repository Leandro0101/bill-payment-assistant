package main

import (
	"bill-payment-assistant/database"
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

	database.StartDB(database.PostgresConfig{
		Host: os.Getenv("POSTGRES_HOST"),
		Pass: os.Getenv("POSTGRES_PASSWORD"),
		User: os.Getenv("POSTGRES_USER"),
		Port: os.Getenv("POSTGRES_PORT"),
		DB: os.Getenv("POSTGRES_DB"),
	})
	
	ginserver.Run(config)
}