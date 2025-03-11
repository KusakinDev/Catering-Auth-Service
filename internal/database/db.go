package database

import (
	"log"

	useraccount "github.com/GIT_USER_ID/GIT_REPO_ID/internal/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	connection *gorm.DB
}

func (database *DataBase) InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=5121508 dbname=catering_auth_db port=5432 sslmode=disable"
	database.connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func (database *DataBase) Migration() {
	database.connection.AutoMigrate(&useraccount.UserAccount{})
}
