package helpers

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SignedDetails struct {
	Username  string
	FirstName string
	LastName  string
	UserType  string
	jwt.RegisteredClaims
}

var SECRET_KEY string = "12345"

func GenerateAllTokens(username string, firstname string, lastname string, usertype string) (signedToken string, err error) {
	claims := &SignedDetails{
		Username:  username,
		FirstName: firstname,
		LastName:  lastname,
		UserType:  usertype,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {

	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)

	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		return
	}

	return claims, msg
}
