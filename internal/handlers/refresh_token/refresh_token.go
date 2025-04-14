package refreshtoken

import (
	jwtconfig "github.com/KusakinDev/Catering-Auth-Service/internal/config/jwt"

	useraccount "github.com/KusakinDev/Catering-Auth-Service/internal/models/account_model"
	jwtt "github.com/KusakinDev/Catering-Auth-Service/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RefreshTokenHandle(c *gin.Context) (int, string, string, string) {

	var reqBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil || reqBody.RefreshToken == "" {
		return 400, "", "", "Refresh token is required"
	}

	// Проверяем валидность refresh токена
	token, err := jwt.Parse(reqBody.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return jwtconfig.JWT_KEY, nil
	})
	if err != nil || !token.Valid {
		return 406, "", "", "Invalid refresh token"
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || token.Method != jwt.SigningMethodHS256 {
		return 406, "", "", "Invalid token claims"
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return 406, "", "", "Invalid token claims"
	}

	var user useraccount.UserAccount
	user.Id = int(userID)
	codeA, accessToken, errAT := jwtt.GenerateAccessToken(user)
	if codeA != 200 {
		return codeA, "", "", errAT
	}

	codeR, refreshToken, errRT := jwtt.GenerateRefreshToken(user)
	if codeR != 200 {
		return codeR, "", "", errRT
	}

	return 200, accessToken, refreshToken, "Token generation is successful"
}
