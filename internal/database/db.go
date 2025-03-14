package database

import (
	"log"

	resetpasswordcode "github.com/KusakinDev/Catering-Auth-Service/internal/models/reset_password_code"
	useraccount "github.com/KusakinDev/Catering-Auth-Service/internal/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	Connection *gorm.DB
}

func (database *DataBase) InitDB() error {
	var err error
	dsn := "host=localhost user=postgres password=5121508 dbname=catering_auth_db port=5432 sslmode=disable"
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

func (database *DataBase) Migration() {
	database.Connection.AutoMigrate(&useraccount.UserAccount{})
	database.Connection.AutoMigrate(&resetpasswordcode.ResetCode{})
}

func (db *DataBase) Create(value interface{}) error {
	return db.Connection.Create(value).Error
}

func (db *DataBase) First(out interface{}, where ...interface{}) error {
	return db.Connection.First(out, where...).Error
}

func (db *DataBase) Save(value interface{}) error {
	return db.Connection.Save(value).Error
}

func (db *DataBase) Find(out interface{}, where ...interface{}) error {
	return db.Connection.Find(out, where...).Error
}
