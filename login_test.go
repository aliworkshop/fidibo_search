package main

import (
	"bytes"
	"fidibo/app"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestLogin(t *testing.T) {
	r := SetUpRouter()
	a := app.New(config())
	a.Init()
	a.InitModules()
	r.POST("/login", a.UserModule.Login)

	// test login with not exists user
	req, _ := http.NewRequest(http.MethodPost, "/login",
		bytes.NewBuffer([]byte(`{"username":"test","password":"123456"}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
