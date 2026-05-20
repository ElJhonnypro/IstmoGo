package utils

import (
	"math"
	"mime/multipart"
	"os"
	"strconv"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

func HasSpecialChar(str string) bool {
	for _, r := range str {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

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
	secret := os.Getenv("JWT_SECRET")
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func HasSymbol(str string) bool {
	for _, r := range str {
		if HasSpecialChar(string(r)) {
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
	if len(phone) < 7 {
		return false
	}

	for i, r := range phone {
		if i == 0 && r == '+' {
			continue
		}
		if !unicode.IsDigit(r) {
			print(r)
			return false
		}
	}

	return true
}

func IsUserName(username string) bool {
	if IsEmail(username) || IsPhoneNumber(username) {
		return false
	}

	for _, r := range username {
		if HasSpecialChar(string(r)) {
			return false
		}
	}

	return true
}

var allowedMIMEs = map[string]bool{
	"image/png":  true,
	"image/jpeg": true,
	"image/jpg":  true,
}

func IsvalidImage(file *multipart.FileHeader) bool {
	mime := file.Header.Get("Content-Type")
	const maxImageSize = 5 * 1024 * 1024 // 5MB

	if file.Size > maxImageSize {
		return false
	}
	return allowedMIMEs[mime]
}

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // km
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func CalculatePrice(distanceKm float64) float64 {
	baseFare := 2.00
	pricePerKm := 0.75

	return baseFare + (distanceKm * pricePerKm)
}

func IsWithinRadius(lat1, lon1, lat2, lon2, radiusKm float64) bool {
	distance := Haversine(lat1, lon1, lat2, lon2)
	return distance <= radiusKm
}

func StringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}
