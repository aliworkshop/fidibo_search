package delivery

import (
	"fidibo/helper"
	"fidibo/user/domain"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type registerHandler struct {
	uc domain.UserUc
}

func NewRegisterHandler(uc domain.UserUc) gin.HandlerFunc {
	h := &registerHandler{
		uc: uc,
	}
	return h.Register
}

func (h *registerHandler) Register(c *gin.Context) {
	var request domain.RegisterRequest
	if err := c.Bind(&request); err != nil {
		helper.ErrorResponse(c, helper.BadRequest, err)
		return
	}

	u := &domain.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Username:  request.Username,
		Active:    true,
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error on generate password hash: ", err.Error())
		helper.ErrorResponse(c, helper.InternalServerError)
	}
	u.Password = string(password)

	err = h.uc.Create(u)
	if err != nil {
		log.Println("error on create new user: ", err.Error())
		helper.ErrorResponse(c, helper.InternalServerError)
	}

	helper.SuccessResponse(c, map[string]any{})
}
