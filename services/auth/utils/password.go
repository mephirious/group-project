package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}

var (
	ErrInvalidPasswordLength = fmt.Errorf("error: invalid password length")
)
