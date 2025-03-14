package registration

import (
	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
	useraccount "github.com/KusakinDev/Catering-Auth-Service/internal/models/user"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

// User's register
func RegistrationHandle(db *database.DataBase, c *gin.Context) (int, string) {

	var user useraccount.UserAccount

	user.DecodeFromContext(c)
	user.SetPasswordHash(user.Password)

	if user.Username == "" || user.Password == "" {
		logger.Errorln("Field is empty")
		return 400, "Field is empty"
	}

	err := user.AddToTable(db)
	if err == 409 {
		logger.Errorln("With user is already exist")
		return 400, "With user is already exist"
	}

	logger.Infoln("Registration of new user is successful. User: ", user.Id, user.Username)
	return 200, "Registration of new user is successful"

}
