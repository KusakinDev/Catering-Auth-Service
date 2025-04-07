package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	Connection *gorm.DB
}

func (database *DataBase) InitDB() error {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	database.Connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error open DB connection: ", err)
		return err
	}
	return nil
}

func (database *DataBase) CloseDB() error {
	sqlDB, err := database.Connection.DB()
	if err != nil {
		log.Println("Error close DB connection: ", err)
		return err
	}
	sqlDB.Close()
	return nil
}
