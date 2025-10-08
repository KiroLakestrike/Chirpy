package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(r *http.Request) (string, error) {
	// get the Bearer information from the http header
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("authorization header not found")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("authorization header not bearer")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	token = strings.TrimSpace(token)

	if token == "" {
		return "", errors.New("authorization header bearer token not found")
	}
	return token, nil
}
