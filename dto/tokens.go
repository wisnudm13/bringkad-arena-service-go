package dto

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	ID   uint   `json:"id"`
	UUID string `json:"uuid"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}
