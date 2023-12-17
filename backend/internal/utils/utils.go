package utils

import (
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func LeftPad(input string, length int, padChar byte) string {
	padLength := length - len(input)
	if padLength <= 0 {
		return input
	}
	padding := strings.Repeat(string(padChar), padLength)
	return padding + input
}

func GenerateAccountNumber() string {
	number := rand.Intn(10000000)

	return LeftPad(fmt.Sprint(number), 8, '0')
}

func EncryptPassword(password string) (string, error) {
    pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(pw), err
}
