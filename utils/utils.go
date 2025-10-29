package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(passwordString string) (string, error) {
	cost := bcrypt.DefaultCost
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(passwordString), cost)
	if err != nil {
		return "", err
	}
	hashedPassword := string(passwordBytes)
	return hashedPassword, nil
}

func IsPasswordValid(passwordFromDB string, passwordFromRequest string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(passwordFromRequest))
	return err == nil
}
