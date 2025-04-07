package rolemodel

import (
	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Role struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`
	Role        string `gorm:"type:varchar(5)"`
	Role_string string `gorm:"type:varchar(20)"`
}

func (role *Role) DecodeFromContext(c *gin.Context) error {

	if err := c.ShouldBindJSON(&role); err != nil {
		logrus.Error("Error decode JSON: ", err)
		return err
	}
	return nil
}

func (role *Role) MigrateToDB(db database.DataBase) error {
	err := db.Connection.AutoMigrate(&Role{})
	if err != nil {
		logrus.Errorln("Error migrate Role model :")
		return err
	}
	logrus.Infoln("Success migrate Role model :")
	return nil
}
