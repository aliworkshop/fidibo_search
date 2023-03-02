package main

import (
	"bytes"
	"encoding/json"
	"fidibo/app"
	"fidibo/user/domain"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearch(t *testing.T) {
	r := SetUpRouter()
	a := app.New(config())
	a.Init()
	a.InitServices()
	a.InitModules()
	r.POST("/search/book", a.MustAuthenticate, a.BookModule.SearchBook)
	r.POST("/login", a.UserModule.Login)

	loginReq := domain.LoginRequest{
		Username: "jesus",
		Password: "123456",
	}
	jsonValue, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	response := map[string]any{}
	json.Unmarshal(responseData, &response)

	assert.Equal(t, http.StatusOK, w.Code)

	// search for book without token
	req, _ = http.NewRequest(http.MethodPost, "/search/book?keyword=test", nil)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// search for book with token
	req, _ = http.NewRequest(http.MethodPost, "/search/book?keyword=test", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+response["access_token"].(string))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
