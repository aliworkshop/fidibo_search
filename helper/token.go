package helper

import (
	"errors"
	"net/http"
	"strings"
)

const (
	authorizationPrefix    = "bearer "
	authorizationPrefixLen = 7

	cookieKeyNormal = "ACCESS_TOKEN_NORMAL"
)

func GetAccessToken(r *http.Request) (string, error) {
	auth := r.Header.Get("authorization")
	if auth != "" {
		if len(auth) < authorizationPrefixLen {
			return "", errors.New("UnAuthorized")
		}
		if strings.ToLower(auth[:authorizationPrefixLen]) != authorizationPrefix {
			return "", errors.New("UnAuthorized")
		}
		token := auth[authorizationPrefixLen:]
		return token, nil
	}
	token := r.FormValue("access_token")
	if token != "" {
		return token, nil
	}
	cookie, _ := r.Cookie(cookieKeyNormal)
	if cookie != nil {
		token = cookie.Value
		if token != "" {
			return token, nil
		}
	}
	return "", errors.New("UnAuthorized")
}
