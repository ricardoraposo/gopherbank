package utils

import (
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func leftPad(input string, length int, padChar byte) string {
	padLength := length - len(input)
	if padLength <= 0 {
		return input
	}
	padding := strings.Repeat(string(padChar), padLength)
	return padding + input
}

func GenerateAccountNumber() string {
	number := rand.Intn(10000000)

	return leftPad(fmt.Sprint(number), 8, '0')
}

func EncryptPassword(password string) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pw), err
}

func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
