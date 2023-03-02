package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTClaim struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	jwt.RegisteredClaims
}

type JwtConfig struct {
	SecretKey            string
	AccessExpirationTime time.Duration
}
