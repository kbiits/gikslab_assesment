package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (hashed string, err error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	hashed = string(hashedBytes)
	return
}
