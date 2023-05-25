package db

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connect()(*gorm.DB, error){
	err:= godotenv.Load(".env")

	if err != nil{
		log.Fatal("Error loading .env file")
	}
	dbUrl := os.Getenv("DB_CONNECTION_STRING")
	conn,err := gorm.Open(postgres.Open(dbUrl),&gorm.Config{})

	if err !=nil {
		return nil,err
	}
	return conn,nil
}
