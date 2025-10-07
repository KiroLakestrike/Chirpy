package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

func GetBearerToken(r *http.Request) (string, error) {
	// get the Bearer information from the http header
	authHeader := r.Header.Get("Authorization")
	log.Printf("DEBUG GetBearerToken: Authorization header = '%s'", authHeader)

	if authHeader == "" {
		return "", errors.New("authorization header not found")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("authorization header not bearer")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	token = strings.TrimSpace(token)
	log.Printf("DEBUG GetBearerToken: Extracted token = '%s'", token)

	if token == "" {
		return "", errors.New("authorization header bearer token not found")
	}
	return token, nil
}
