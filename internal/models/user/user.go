package useraccount

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserAccount struct {
	// User account id in UserAccount table
	id int `gorm:"primaryKey;autoIncrement"`

	// User account username in UserAccount table
	username string `gorm:"type:varchar(20)"`

	// User account password in UserAccount table
	password string `gorm:"type:varchar(100)"`
}

// Get user's Id
func (user *UserAccount) GetId() int {
	return user.id
}

// Set user's Id
func (user *UserAccount) SetId(idIN int) {
	user.id = idIN
}

// Get user's username
func (user *UserAccount) GetUsername() string {
	return user.username
}

// Set user's username
func (user *UserAccount) SetUsername(usernameIN string) {
	user.username = usernameIN
}

// Get user's password as hash
func (user *UserAccount) GetPassword() string {
	return user.password
}

// Set user's password as hash (in:password out:hash)
func (user *UserAccount) SetPassword(passwordIN string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordIN), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.password = string(hashedPassword)
	return nil
}

// Decode struct from json gin context
func (user *UserAccount) DecodeFromContext(c *gin.Context) error {
	if err := c.ShouldBindJSON(&user); err != nil {
		return err
	}
	return nil
}
