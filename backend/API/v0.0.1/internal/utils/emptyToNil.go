package utils

import (
	"os"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v5"
)

func EmptyToNil(s *string) *string {
	if s != nil && *s == "" {
		return nil
	}
	return s
}

func NilToEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(), // 7 días
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(os.Getenv("JWT_SECRET"))
}

func HasSymbol(str string) bool {
	for _, r := range str {
		if unicode.IsSymbol(r) {
			return true
		}
	}
	return false
}

func IsEmail(email string) bool {
	for _, r := range email {
		if r == '@' {
			return true
		}
	}
	return false
}

func IsPhoneNumber(phone string) bool {
	for _, r := range phone {
		if !unicode.IsDigit(r) {
			return false
		}

		if r == '+' {
			continue
		}
	}
	return true
}

func IsUserName(username string) bool {
	if IsEmail(username) || IsPhoneNumber(username) {
		return false
	}

	for _, r := range username {
		if unicode.IsSymbol(r) {
			return false
		}
	}

	return true
}
