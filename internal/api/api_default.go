package api

import (
	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
	"github.com/KusakinDev/Catering-Auth-Service/internal/handlers/login"
	refreshtoken "github.com/KusakinDev/Catering-Auth-Service/internal/handlers/refreshToken"
	"github.com/KusakinDev/Catering-Auth-Service/internal/handlers/registration"
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

	code, accessToken, refreshToken, message := refreshtoken.RefreshToken(c)
	c.JSON(code, gin.H{"accessToken": accessToken, "refreshToken": refreshToken, "message": message})
}
