package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
	}

	return err
}

func MatchUserTypeToUsername(c *gin.Context, username string) (err error) {
	userType := c.GetString("user_type")
	uName := c.GetString("username")
	err = nil

	if userType == "USER" && uName != username {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	return
}
