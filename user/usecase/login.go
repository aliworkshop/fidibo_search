package usecase

import (
	"fidibo/helper"
	"fidibo/user/domain"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func (uc *usecase) Login(user *domain.User) (string, int, error) {
	claims := helper.JWTClaim{
		Name:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        fmt.Sprintf("%d", user.ID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(uc.config.AccessExpirationTime)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(uc.config.SecretKey))
	if err != nil {
		return "", 0, err
	}

	return t, int(uc.config.AccessExpirationTime.Seconds()), nil
}
