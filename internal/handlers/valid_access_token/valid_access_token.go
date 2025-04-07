package validaccesstoken

import (
	"net/http"

	jwtconfig "github.com/KusakinDev/Catering-Auth-Service/internal/config/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidAccessToken(c *gin.Context) (int, jwt.MapClaims) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return 401, jwt.MapClaims{}
	}

	tokenString = tokenString[len("Bearer "):]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtconfig.JWT_KEY, nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return 401, jwt.MapClaims{}
	}
	return 200, token.Claims.(jwt.MapClaims)
}
