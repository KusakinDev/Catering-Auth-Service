package useraccount

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DBInterface interface {
	Create(value interface{}) error
	First(out interface{}, where ...interface{}) error
	Save(value interface{}) error
	Find(out interface{}, where ...interface{}) error
}

type UserAccount struct {
	// User account id in UserAccount table
	Id int `gorm:"primaryKey;autoIncrement"`

	// User account username in UserAccount table
	Username string `gorm:"type:varchar(20)"`

	// User account email in UserAccount table
	Email string `gorm:"type:varchar(50)"`

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
func (user *UserAccount) AddToTable(db DBInterface) int {
	userFind := &UserAccount{}
	userFind.Username = user.Username

	errFind := userFind.GetFromTableByName(db)
	if errFind == nil {
		logrus.Println("userFind ", userFind)
		return 409
	}

	err := db.Create(&user)
	if err != nil {
		return 503
	}
	return 0
}

// Get user from table by username
func (user *UserAccount) GetFromTableByName(db DBInterface) error {
	err := db.First(&user, "username = ?", user.Username)
	if err != nil {
		return err
	}
	return nil
}

// Get user from table by id
func (user *UserAccount) GetFromTable(db DBInterface) error {
	err := db.First(&user)
	if err != nil {
		return err
	}
	return nil
}

// Update user in table by id
func (user *UserAccount) UpdateInTable(db DBInterface) error {
	err := db.Save(&user)
	if err != nil {
		return err
	}
	return nil
}
