package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Muhammad-Mahir157/clockify-app-clone/config"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &config.Config{
		Host:     os.Getenv("DB_Host"),
		Port:     os.Getenv("DB_Port"),
		User:     os.Getenv("DB_User"),
		Password: os.Getenv("DB_Password"),
		DBName:   os.Getenv("DB_Name"),
		SSLMode:  os.Getenv("DB_Mode"),
	}

	db, err := NewConnection(config)
	if err != nil {
		log.Fatal("Could not load the database ...")
	}

	Db = db

}

func NewConnection(con *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s password=%s user=%s dbname=%s sslmode=%s",
		con.Host, con.Port, con.Password, con.User, con.DBName, con.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
