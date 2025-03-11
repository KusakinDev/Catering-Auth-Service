package api

import (
	"github.com/KusakinDev/Catering-Auth-Service/internal/database"
	"github.com/KusakinDev/Catering-Auth-Service/internal/handlers/login"
	"github.com/gin-gonic/gin"
)

type DefaultAPI struct {
	db database.DataBase
}

// Post /register
// New user's registration
func (api *DefaultAPI) RegisterPost(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Post /login
// User's login
func (api *DefaultAPI) Login(c *gin.Context) {

	err := api.db.InitDB()
	if err != nil {
		c.JSON(503, gin.H{"Error": "Service Unavailable"})
	}
	defer api.db.CloseDB()

	code, message := login.LoginHandle(&api.db, c)

	c.JSON(code, gin.H{"message": message})

}

// Post /refresh-token
// Refresh access token
func (api *DefaultAPI) RefreshTokenPost(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}
