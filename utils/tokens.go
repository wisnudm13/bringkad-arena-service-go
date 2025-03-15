package utils

import (
	"bringkad-arena-service-go/dto"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(accountType string, uuid string, id uint) string {
	var accessTokenSecret = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	claims := dto.JWTClaim{
		ID:   id,
		UUID: uuid,
		Type: accountType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(accessTokenSecret)

	return token
}

func GenerateRefreshAccessToken(accountType string, uuid string, id uint) string {
	var refreshTokenSecret = []byte(os.Getenv("REFRESH_ACCESS_TOKEN_SECRET"))
	claims := dto.JWTClaim{
		ID:   id,
		UUID: uuid,
		Type: accountType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(refreshTokenSecret)

	return token
}

func DecodeJWT(JWTtoken string, JWTSecret []byte) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(JWTtoken, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, fmt.Errorf("Unauthorized")

}
