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

// Get user's Id
func (user *UserAccount) GetId() int {
	return user.Id
}

// Set user's Id
func (user *UserAccount) SetId(idIN int) {
	user.Id = idIN
}

// Get user's username
func (user *UserAccount) GetUsername() string {
	return user.Username
}

// Set user's username
func (user *UserAccount) SetUsername(usernameIN string) {
	user.Username = usernameIN
}

// Get user's password as hash
func (user *UserAccount) GetPassword() string {
	return user.Password
}

func (user *UserAccount) SetPassword(passwordIN string) {
	hasher := sha256.New()
	hasher.Write([]byte(passwordIN))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	user.Password = hashedPassword
}

// Decode struct from json gin context
func (user *UserAccount) DecodeFromContext(c *gin.Context) error {
	var temp struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&temp); err != nil {
		logrus.Error("Error decode JSON: ", err)
		return err
	}

	user.Id = temp.ID
	user.Username = temp.Username
	user.Password = temp.Password

	return nil
}

func (user *UserAccount) AddToTable(db *database.DataBase) error {
	return nil
}

// Get user from table by id or username
func (user *UserAccount) GetFromTable(db *database.DataBase) error {

	err := db.Connection.First(&user).Error
	if err != nil {
		return err
	}

	return nil
}
