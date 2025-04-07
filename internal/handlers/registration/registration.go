package registration

import (
	useraccount "github.com/KusakinDev/Catering-Auth-Service/internal/models/account_model"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

// User's register
func RegistrationHandle(c *gin.Context) (int, string) {

	var user useraccount.UserAccount

	user.DecodeFromContext(c)
	user.SetPasswordHash(user.Password)

	if user.Email == "" || user.Password == "" {
		logger.Errorln("Field is empty")
		return 400, "Field is empty"
	}

	err := user.AddToTable()
	if err == 409 {
		logger.Errorln("With user is already exist")
		return 400, "With user is already exist"
	}
	if err == 503 {
		logger.Errorln("Not avalible")
		return 503, "Not avalible"
	}

	logger.Infoln("Registration of new user is successful. User: ", user.Id, user.Email)
	return 200, "Registration of new user is successful"

}
