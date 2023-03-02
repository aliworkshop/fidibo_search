package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error string

func (e Error) String() string {
	return string(e)
}

const (
	BadRequest          Error = "Bad Request"
	InvalidCredentials  Error = "Invalid Credentials"
	InternalServerError Error = "Internal Server Error"
	UnAuthorized        Error = "UnAuthorized"
)

var Errors = map[Error]int{
	BadRequest:          http.StatusBadRequest,
	InvalidCredentials:  http.StatusUnauthorized,
	InternalServerError: http.StatusInternalServerError,
	UnAuthorized:        http.StatusUnauthorized,
}

func ErrorResponse(c *gin.Context, e Error, err ...error) {
	data := map[string]any{
		"success": false,
		"message": e.String(),
	}
	if len(err) > 0 {
		data["error"] = err[0].Error()
	}
	c.JSON(Errors[e], data)
	c.Abort()
}

func SuccessResponse(c *gin.Context, data map[string]any) {
	data["success"] = true
	c.JSON(http.StatusOK, data)
}
