package resetpasswordcode

import (
	"crypto/rand"
	"math/big"
	"time"

	useraccount "github.com/KusakinDev/Catering-Auth-Service/internal/models/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DBInterface interface {
	Create(value interface{}) error
	First(out interface{}, where ...interface{}) error
	Save(value interface{}) error
	Find(out interface{}, where ...interface{}) error
	Delete(value interface{}, where ...interface{}) error
}

type ResetCode struct {
	Id        int                     `gorm:"primaryKey;autoIncrement"`
	Id_user   int                     `gorm:"not null"`
	Code      int                     `gorm:"not null;type:integer"`
	StartTime string                  `gorm:"type:varchar(50)"`
	ExpTime   string                  `gorm:"type:varchar(50)"`
	User      useraccount.UserAccount `gorm:"foreignKey:Id_user;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Decode struct from json gin context
func (resetForm *ResetCode) DecodeFromContext(c *gin.Context) error {

	if err := c.ShouldBindJSON(&resetForm); err != nil {
		logrus.Error("Error decode JSON: ", err)
		return err
	}
	return nil
}

// Generate reset code
func (resetForm *ResetCode) GenerateCode() error {
	min := int64(100000)
	max := int64(999999)

	randNum, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		logrus.Error("Error generate reset code: ", err)
		return err
	}

	resetForm.Code = int(randNum.Int64() + min)

	return nil
}

// Init date for code, where param is duration in mitune
func (resetForm *ResetCode) InitDate(duration int) {

	currentTime := time.Now()
	expirationTime := currentTime.Add(time.Duration(duration) * time.Minute)

	resetForm.StartTime = currentTime.Format("2006-01-02 15:04:05")
	resetForm.ExpTime = expirationTime.Format("2006-01-02 15:04:05")
}

// Create new row in code
func (resetForm *ResetCode) AddToTable(db DBInterface) int {
	resetForm.Code = int(resetForm.Code)
	err := db.Create(&resetForm)
	if err != nil {
		return 503
	}
	return 0
}

// Get coderecord from table by code
func (resetForm *ResetCode) GetFromTableByCode(db DBInterface) error {
	err := db.First(&resetForm, "code = ?", int(resetForm.Code))
	if err != nil {
		return err
	}
	return nil
}

// Get coderecord from table by user id
func (resetForm *ResetCode) GetFromTableByUserId(db DBInterface) error {
	err := db.First(&resetForm, "id_user = ?", resetForm.Id_user)
	if err != nil {
		return err
	}
	return nil
}

// Remove coderecord from table by code
func (resetForm *ResetCode) DeleteFromTableByCode(db DBInterface) error {
	err := db.Delete(&resetForm, "code = ?", resetForm.Code)
	if err != nil {
		logrus.Error("Error deleting reset code: ", err)
		return err
	}
	return nil
}

func (resetForm *ResetCode) ValideCode() (int, string) {
	dateExp, _ := time.Parse("2006-01-02 15:04:05", resetForm.ExpTime)
	dateNow := time.Now()
	if dateExp.Before(dateNow) {
		return 403, "Code is not valide"
	}
	return 200, ""
}
