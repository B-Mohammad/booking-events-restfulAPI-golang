package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	hashedP, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	return string(hashedP), err
}

func ComparePassword(plainPass string, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass))
	return err == nil
}
