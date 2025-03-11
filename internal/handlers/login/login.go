package login

import (
	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
	useraccount "github.com/KusakinDev/Catering-Auth-Service/internal/models/user"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

// User's login
func LoginHandle(db *database.DataBase, c *gin.Context) (int, string) {

	var userFront useraccount.UserAccount
	var userDB useraccount.UserAccount

	userFront.DecodeFromContext(c)
	userFront.SetPassword(userFront.GetPassword())

	userDB.SetUsername(userFront.GetUsername())

	err := userDB.GetFromTable(db)
	if err != nil {
		logger.Errorln("Incorrect login")
		return 403, "Incorrect login or password"
	}

	if userDB.GetPassword() != userFront.GetPassword() {
		logger.Errorln("Incorrect password")
		return 403, "Incorrect login or password"
	}

	logger.Infoln("Authorization is successful. User: ", userDB.Id, userDB.GetUsername())
	return 200, "Authorization is successful"

}
