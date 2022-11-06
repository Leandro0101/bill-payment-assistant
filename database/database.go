package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host string
	Pass string
	User string
	Port string
	DB string
}

var db *gorm.DB

func StartDB(config PostgresConfig){
	dsn := "host=" + config.Host + " user=" + config.User + " password=" + config.Pass + " dbname=" + config.DB + " port=" + config.Port + " sslmode=disable TimeZone=UTC"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal(err)
	}

	db = database
}

func GetDatabase() *gorm.DB {
	return db
}