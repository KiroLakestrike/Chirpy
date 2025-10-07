package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

// GetRefreshToken extracts the refresh token from the Authorization header
// Refresh tokens can be sent with or without "Bearer " prefix
func GetRefreshToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	log.Printf("DEBUG GetRefreshToken: Authorization header = '%s'", authHeader)

	if authHeader == "" {
		return "", errors.New("authorization header not found")
	}

	// Try to extract token with Bearer prefix first
	if strings.HasPrefix(authHeader, "Bearer ") {
		token := strings.TrimPrefix(authHeader, "Bearer ")
		token = strings.TrimSpace(token)
		log.Printf("DEBUG GetRefreshToken: Extracted token with Bearer = '%s'", token)
		if token == "" {
			return "", errors.New("token not found")
		}
		return token, nil
	}

	// If no Bearer prefix, return the header value directly (trimmed)
	token := strings.TrimSpace(authHeader)
	log.Printf("DEBUG GetRefreshToken: Extracted token without Bearer = '%s'", token)
	if token == "" {
		return "", errors.New("token not found")
	}
	return token, nil
}
