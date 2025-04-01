package changepassword

import (
	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
	changepasswordform "github.com/KusakinDev/Catering-Auth-Service/internal/models/change_password_form"
	useraccount "github.com/KusakinDev/Catering-Auth-Service/internal/models/user"
	"github.com/KusakinDev/Catering-Auth-Service/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

// User's login
func ChangePasswordHandle(db *database.DataBase, c *gin.Context) (int, string, string, string) {

	var user useraccount.UserAccount
	var form changepasswordform.ChangePasswordForm
	var err error

	err = form.DecodeFromContext(c)
	if err != nil {
		return 400, "", "", "No data input"
	}

	form.HashPassword()

	if form.NewPassword == form.OldPassword {
		return 400, "", "", "Passwords mustn't match"
	}

	user.Username = form.Username
	err = user.GetFromTableByName(db)
	if err != nil {
		logger.Errorln("Incorrect login")
		return 403, "", "", "Incorrect login or password"
	}

	if user.Password != form.OldPassword {
		logger.Errorln("Incorrect password")
		return 403, "", "", "Incorrect login or password"
	}

	user.Password = form.NewPassword

	user.UpdateInTable(db)

	codeA, accessToken, errAT := jwt.GenerateAccessToken(user)
	if codeA != 200 {
		return codeA, "", "", errAT
	}

	codeR, refreshToken, errRT := jwt.GenerateRefreshToken(user)
	if codeR != 200 {
		return codeR, "", "", errRT
	}

	logger.Infoln("Change password for user is successful. User: ", user.Id, user.Username)

	return 200, accessToken, refreshToken, "Change password is successful"
}
