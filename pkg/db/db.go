package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupConn() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file, doesn't exist.")
	}

	DB_URL := os.Getenv("DB_URL")

	conn, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatal("Caannot connect to the database.")
	}
	log.Println("Connected to database.")
	return conn
}
