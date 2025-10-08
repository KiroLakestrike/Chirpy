package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(r *http.Request) (string, error) {
	// get the API Key from the HEader
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("authorization header not found")
	}

	if !strings.HasPrefix(authHeader, "ApiKey ") {
		return "", errors.New("authorization API key not found")
	}

	token := strings.TrimPrefix(authHeader, "ApiKey ")
	token = strings.TrimSpace(token)

	if token == "" {
		return "", errors.New("authorization Api key not found")
	}
	return token, nil
}
