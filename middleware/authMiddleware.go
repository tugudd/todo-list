package middleware

import (
	"fmt"
	"main/helpers"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			c.JSON(500, gin.H{"error": fmt.Sprintf("No authorization header provided")})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		fmt.Println(777)

		if err != "" {
			c.JSON(500, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("first_name", claims.FirstName)
		c.Set("last_name", claims.LastName)
		c.Set("user_type", claims.UserType)
		c.Next()
	}
}
