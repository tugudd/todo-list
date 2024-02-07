package service

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashedPass)
}

func VerifyPassword(providedPassword string, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))

	if err != nil {
		return false
	}

	return true
}
