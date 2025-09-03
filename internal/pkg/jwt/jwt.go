package jwt

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(ID uint, email string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	defaultHourExp := os.Getenv("JWT_EXPIRED")
	convertInt, _ := strconv.Atoi(defaultHourExp)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    ID,
		"exp":   time.Now().Add(time.Hour * time.Duration(convertInt)),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
