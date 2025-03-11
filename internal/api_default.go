package api

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/internal/database"
	"github.com/gin-gonic/gin"
)

type DefaultAPI struct {
	db database.DataBase
}

// Post /login
// User's login
func (api *DefaultAPI) LoginPost(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Post /refresh-token
// Refresh access token
func (api *DefaultAPI) RefreshTokenPost(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Post /register
// New user's registration
func (api *DefaultAPI) RegisterPost(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}
