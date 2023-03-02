package main

import (
	"bytes"
	"encoding/json"
	"fidibo/app"
	"fidibo/user/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	r := SetUpRouter()
	a := app.New(config())
	a.Init()
	a.InitModules()
	r.POST("/register", a.UserModule.Register)
	r.POST("/login", a.UserModule.Login)

	registerReq := domain.RegisterRequest{
		Username:  "jesus",
		Password:  "123456",
		Email:     "jesus@gmail.com",
		FirstName: "jesus",
		LastName:  "",
	}
	jsonValue, _ := json.Marshal(registerReq)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	loginReq := domain.LoginRequest{
		Username: registerReq.Username,
		Password: registerReq.Password,
	}
	jsonValue, _ = json.Marshal(loginReq)
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
