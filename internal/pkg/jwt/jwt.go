package jwt

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	UserID uint
	Email  string
	jwt.RegisteredClaims
}

func GenerateJwt(ID uint, email string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := &JwtClaims{
		UserID: ID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshJwt(ID uint, email string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	expiredHour := os.Getenv("JWT_EXPIRED")
	convertInt, _ := strconv.Atoi(expiredHour)

	claims := &JwtClaims{
		UserID: ID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(convertInt))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJwt(tokenString string) (*JwtClaims, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	claims := &JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
