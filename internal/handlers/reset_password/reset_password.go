package resetpassword

import (
	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
	resetpasswordcode "github.com/KusakinDev/Catering-Auth-Service/internal/models/reset_password_code"
	useraccount "github.com/KusakinDev/Catering-Auth-Service/internal/models/user"
	"github.com/KusakinDev/Catering-Auth-Service/internal/utils/email"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// User's reset password
func ResetPasswordHandle(db *database.DataBase, c *gin.Context) (int, string) {

	var userFront useraccount.UserAccount
	var userDB useraccount.UserAccount
	var err error

	userFront.DecodeFromContext(c)
	userDB.Username = userFront.Username

	err = userDB.GetFromTableByName(db)
	if err != nil {
		logrus.Errorln("Incorrect login")
		return 403, "Incorrect login or email"
	}

	if userDB.Email != userFront.Email {
		logrus.Errorln("Incorrect email")
		return 403, "Incorrect login or email"
	}

	var resetForm resetpasswordcode.ResetCode

	resetForm.GenerateCode()
	resetForm.InitDate(5)

	err = email.SendEmail(userDB.Email, userDB.Username, resetForm.Code)
	if err != nil {
		return 503, "Error send email"
	}

	resetForm.User = userDB
	resetForm.AddToTable(db)

	logrus.Infoln("Send reset email is successful. User: ", userDB.Id, userDB.Username,
		" EMAIL: ", userDB.Email, "code: ", resetForm.Code, "time: ", resetForm.StartTime, " ", resetForm.ExpTime)

	return 200, "Send reset email is successful"
}
