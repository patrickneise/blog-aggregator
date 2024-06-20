package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (token string, err error) {
	authHeader := headers.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("no 'Authorization' header included in request")
	}
	authSplit := strings.Split(authHeader, " ")
	if authSplit[0] != "ApiKey" || len(authSplit) < 2 {
		return "", errors.New("malformed 'Authorization' header")
	}

	return authSplit[1], nil
}
