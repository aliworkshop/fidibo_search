package app

import (
	"fidibo/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (a *App) MustAuthenticate(c *gin.Context) {
	token, err := helper.GetAccessToken(c.Request)
	if err != nil {
		helper.ErrorResponse(c, helper.UnAuthorized)
		return
	}

	claimFunc := func() jwt.Claims {
		return new(helper.JWTClaim)
	}()
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte((a.config.Jwt.SecretKey)), nil
	}
	parsed, e := jwt.ParseWithClaims(token, claimFunc, keyFunc)
	if e != nil {
		helper.ErrorResponse(c, helper.UnAuthorized, e)
		return
	}

	if !parsed.Valid {
		helper.ErrorResponse(c, helper.UnAuthorized)
		return
	}
}
