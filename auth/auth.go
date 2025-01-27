package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string, jwtKey []byte) (string, error) {
	//generating token         setting sign method    additional info stored in token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	//return signed token
	return token.SignedString(jwtKey)
}

// defining middleware. returns gin.HandlerFunc which takes ctx
func AuthMiddleware(jwtKey []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//get token from header
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(401, gin.H{
				"error": "Invalid token",
			})
			ctx.Abort()
			return
		}

		//parse token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		//check token
		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{
				"error": "Invalid token",
			})
			ctx.Abort()
			return
		}

		//go to next function (another middleware or some handler)
		ctx.Next()
	}
}
