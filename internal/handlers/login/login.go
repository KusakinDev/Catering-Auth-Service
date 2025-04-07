package login

import (
	useraccount "github.com/KusakinDev/Catering-Auth-Service/internal/models/account_model"
	"github.com/KusakinDev/Catering-Auth-Service/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

// User's login
func LoginHandle(c *gin.Context) (int, string, string, string) {
	var userFront useraccount.UserAccount
	var userDB useraccount.UserAccount

	userFront.DecodeFromContext(c)
	userFront.SetPasswordHash(userFront.Password)

	userDB.Email = userFront.Email

	err := userDB.GetFromTable()
	if err != nil {
		logger.Errorln("Incorrect login")
		return 403, "", "", "Incorrect login or password"
	}

	if userDB.Password != userFront.Password {
		logger.Errorln("Incorrect password")
		return 403, "", "", "Incorrect login or password"
	}

	codeA, accessToken, errAT := jwt.GenerateAccessToken(userDB)
	if codeA != 200 {
		return codeA, "", "", errAT
	}

	codeR, refreshToken, errRT := jwt.GenerateRefreshToken(userDB)
	if codeR != 200 {
		return codeR, "", "", errRT
	}

	logger.Infoln("Authorization is successful. User: ", userDB.Id, userDB.Email)

	return 200, accessToken, refreshToken, "Authorization is successful"
}
