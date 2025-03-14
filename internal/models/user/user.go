package useraccount

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserAccount struct {
	// User account id in UserAccount table
	Id int `gorm:"primaryKey;autoIncrement"`

	// User account username in UserAccount table
	Username string `gorm:"type:varchar(20)"`

	// User account password in UserAccount table
	Password string `gorm:"type:varchar(100)"`
}

// Set password as hash
func (user *UserAccount) SetPasswordHash(passwordIN string) {
	hasher := sha256.New()
	hasher.Write([]byte(passwordIN))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	user.Password = hashedPassword
}

// Decode struct from json gin context
func (user *UserAccount) DecodeFromContext(c *gin.Context) error {

	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Error("Error decode JSON: ", err)
		return err
	}
	return nil
}

// Create new row in user's account
func (user *UserAccount) AddToTable(db *database.DataBase) int {

	userFind := &UserAccount{}
	userFind.Username = user.Username

	errFind := userFind.getFromTableByName(db)

	if errFind == nil {
		logrus.Println("userFind ", userFind)
		return 409
	}

	err := db.Connection.Create(&user).Error
	if err != nil {
		return 503
	}
	return 0
}

// Get user from table by username
func (user *UserAccount) getFromTableByName(db *database.DataBase) error {

	err := db.Connection.Where("username = ?", user.Username).First(&user).Error
	if err != nil {
		return err
	}

	return nil
}

// Get user from table by id
func (user *UserAccount) GetFromTable(db *database.DataBase) error {

	err := db.Connection.First(&user).Error
	if err != nil {
		return err
	}

	return nil
}
