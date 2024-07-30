package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password []byte) string {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}
