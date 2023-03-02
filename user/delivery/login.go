package delivery

import (
	"fidibo/helper"
	"fidibo/user/domain"
	"github.com/gin-gonic/gin"
	"log"
)

type loginHandler struct {
	uc domain.UserUc
}

func NewLoginDelivery(uc domain.UserUc) gin.HandlerFunc {
	h := &loginHandler{
		uc: uc,
	}
	return h.Login
}

func (h *loginHandler) Login(c *gin.Context) {
	var request domain.LoginRequest
	if err := c.Bind(&request); err != nil {
		helper.ErrorResponse(c, helper.BadRequest, err)
		return
	}

	user, err := h.uc.GetByUsername(request.Username)
	if err != nil {
		helper.ErrorResponse(c, helper.InvalidCredentials)
		return
	}

	ok, err := h.uc.PasswordMatches(request.Password, user)
	if err != nil || !ok {
		helper.ErrorResponse(c, helper.InvalidCredentials)
		return
	}

	token, expires, err := h.uc.Login(user)
	if err != nil {
		log.Println("error on generate access token: ", err.Error())
		helper.ErrorResponse(c, helper.InternalServerError)
		return
	}

	helper.SuccessResponse(c, map[string]any{
		"access_token": token,
		"expires_in":   expires,
	})
}
