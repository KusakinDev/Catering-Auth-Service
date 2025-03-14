package resetpasswordcode

import (
	"crypto/rand"
	"math/big"
	"time"

	useraccount "github.com/KusakinDev/Catering-Auth-Service/internal/models/user"
	"github.com/sirupsen/logrus"
)

type DBInterface interface {
	Create(value interface{}) error
	First(out interface{}, where ...interface{}) error
	Save(value interface{}) error
	Find(out interface{}, where ...interface{}) error
}

type ResetCode struct {
	Id        int                     `gorm:"primaryKey;autoIncrement"`
	Id_user   int                     `gorm:"primaryKey;autoIncrement"`
	Code      int64                   `gorm:"not null"`
	StartTime string                  `gorm:"type:varchar(50)"`
	ExpTime   string                  `gorm:"type:varchar(50)"`
	User      useraccount.UserAccount `gorm:"foreignKey:Id_user;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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

	resetForm.Code = randNum.Int64() + min

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
	err := db.Create(&resetForm)
	if err != nil {
		return 503
	}
	return 0
}
