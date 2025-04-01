package api

import (
	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
	changepassword "github.com/KusakinDev/Catering-Auth-Service/internal/handlers/change_password"
	"github.com/KusakinDev/Catering-Auth-Service/internal/handlers/login"
	refreshtoken "github.com/KusakinDev/Catering-Auth-Service/internal/handlers/refreshToken"
	"github.com/KusakinDev/Catering-Auth-Service/internal/handlers/registration"
	resetpassword "github.com/KusakinDev/Catering-Auth-Service/internal/handlers/reset_password"
	verefyresetcode "github.com/KusakinDev/Catering-Auth-Service/internal/handlers/verefy_reset_code"
	"github.com/gin-gonic/gin"
)

type DefaultAPI struct {
	db database.DataBase
}

// Post /register
// New user's registration
func (api *DefaultAPI) Register(c *gin.Context) {

	err := api.db.InitDB()
	if err != nil {
		c.JSON(503, gin.H{"Error": "Service Unavailable"})
	}
	defer api.db.CloseDB()

	code, message := registration.RegistrationHandle(&api.db, c)

	c.JSON(code, gin.H{"message": message})
}

// Post /login
// User's login
func (api *DefaultAPI) Login(c *gin.Context) {

	err := api.db.InitDB()
	if err != nil {
		c.JSON(503, gin.H{"Error": "Service Unavailable"})
	}
	defer api.db.CloseDB()

	code, accessToken, refreshToken, message := login.LoginHandle(&api.db, c)

	c.JSON(code, gin.H{"accessToken": accessToken, "refreshToken": refreshToken, "message": message})

}

// Post /refresh-token
// Refresh access token
func (api *DefaultAPI) RefreshToken(c *gin.Context) {

	code, accessToken, refreshToken, message := refreshtoken.RefreshTokenHandle(c)
	c.JSON(code, gin.H{"accessToken": accessToken, "refreshToken": refreshToken, "message": message})
}

// Post /change-password
// Change user's password
func (api *DefaultAPI) ChangePassword(c *gin.Context) {

	err := api.db.InitDB()
	if err != nil {
		c.JSON(503, gin.H{"Error": "Service Unavailable"})
	}
	defer api.db.CloseDB()

	code, accessToken, refreshToken, message := changepassword.ChangePasswordHandle(&api.db, c)

	c.JSON(code, gin.H{"accessToken": accessToken, "refreshToken": refreshToken, "message": message})
}

// Post /reset-password
// Reset user's password
func (api *DefaultAPI) ResetPassword(c *gin.Context) {

	err := api.db.InitDB()
	if err != nil {
		c.JSON(503, gin.H{"Error": "Service Unavailable"})
	}
	defer api.db.CloseDB()

	code, message := resetpassword.ResetPasswordHandle(&api.db, c)

	c.JSON(code, gin.H{"message": message})
}

// Post /verefy-reset-code
// Verefy code for reset user's password
func (api *DefaultAPI) VerefyRecetCode(c *gin.Context) {

	err := api.db.InitDB()
	if err != nil {
		c.JSON(503, gin.H{"Error": "Service Unavailable"})
	}
	defer api.db.CloseDB()

	code, accessToken, refreshToken, message := verefyresetcode.VerefyResetCodeHandle(&api.db, c)

	c.JSON(code, gin.H{"accessToken": accessToken, "refreshToken": refreshToken, "message": message})
}
