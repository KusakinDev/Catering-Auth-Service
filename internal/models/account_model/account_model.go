package accountmodel

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
	rolemodel "github.com/KusakinDev/Catering-Auth-Service/internal/models/role_model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserAccount struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(100)"`
	RoleId   int    `json:"role_id" gorm:"not null"`

	Role rolemodel.Role `gorm:"foreignKey:RoleId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
func (user *UserAccount) AddToTable() int {
	var db database.DataBase
	db.InitDB()

	userFind := &UserAccount{}
	userFind.Email = user.Email

	errFind := userFind.GetFromTableByEmail()
	if errFind == nil {
		db.CloseDB()
		logrus.Println("userFind ", userFind)
		return 409
	}

	err := db.Connection.Create(&user).Error
	if err != nil {
		logrus.Println(err)
		logrus.Println(user)
		db.CloseDB()
		return 503
	}
	db.CloseDB()
	return 0
}

// Get user from table by username
func (user *UserAccount) GetFromTableByEmail() error {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.First(&user, "email = ?", user.Email).Error
	if err != nil {
		db.CloseDB()
		return err
	}
	db.CloseDB()
	return nil
}

// Get user from table by id
func (user *UserAccount) GetFromTable() error {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.First(&user).Error
	if err != nil {
		db.CloseDB()
		return err
	}
	db.CloseDB()
	return nil
}

// Update user in table by id
func (user *UserAccount) UpdateInTable() error {
	var db database.DataBase
	db.InitDB()

	err := db.Connection.Save(&user).Error
	if err != nil {
		db.CloseDB()
		return err
	}
	db.CloseDB()
	return nil
}

func (user *UserAccount) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&UserAccount{})
	if err != nil {
		logrus.Errorln("Error migrate UserAccount model :")
		return err
	}
	logrus.Infoln("Success migrate UserAccount model :")
	return nil
}
