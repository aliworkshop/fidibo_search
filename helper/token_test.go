package helper

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestEmptyAccessToken(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, "/search/book", nil)
	req.Header.Set("Content-Type", "application/json")

	token, err := GetAccessToken(req)
	assert.Equal(t, errors.New("UnAuthorized"), err)
	assert.Equal(t, "", token)
}

func TestNotEmptyAccessToken(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, "/search/book", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoi2LnZhNuMINiq2LHYp9io24wiLCJlbWFpbCI6InNyYWxpdG9yYWJpQGdtYWlsLmNvbSIsImV4cCI6MTY3Nzc0NjE4MiwianRpIjoiMSJ9.XpHc2z8YVFVojDEDdSBDCjz4WZhyfQfd1FffjlIx08I")

	token, err := GetAccessToken(req)
	assert.Equal(t, nil, err)
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoi2LnZhNuMINiq2LHYp9io24wiLCJlbWFpbCI6InNyYWxpdG9yYWJpQGdtYWlsLmNvbSIsImV4cCI6MTY3Nzc0NjE4MiwianRpIjoiMSJ9.XpHc2z8YVFVojDEDdSBDCjz4WZhyfQfd1FffjlIx08I", token)
}
