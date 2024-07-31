package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password []byte) string {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func ComparePassword(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return false
	}
	return true
}
