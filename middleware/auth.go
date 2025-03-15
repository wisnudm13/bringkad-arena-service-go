package middleware

import (
	"bringkad-arena-service-go/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func ValidateAccessToken(accountType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var accessToken string
		var err error

		if accountType == "admin" {
			accessToken, err = c.Cookie("access_token")

			if err != nil {
				utils.UnauthorizedMessage(c)
				c.Abort()
				return
			}

		} else {
			// tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
			// token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// 	return jwtSecret, nil
			// })

			// if err != nil || !token.Valid {

			// }
		}

		claims, err := utils.DecodeJWT(accessToken, []byte(os.Getenv("ACCESS_TOKEN_SECRET")))

		if err != nil {
			utils.UnauthorizedMessage(c)
			c.Abort()
			return
		}

		fmt.Println(claims)

		c.Next()
	}
}
