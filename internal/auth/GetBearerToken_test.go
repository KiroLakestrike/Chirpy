package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name          string
		authHeader    string
		expectedToken string
		expectError   bool
	}{
		{
			name:          "Valid Bearer token",
			authHeader:    "Bearer abc123token",
			expectedToken: "abc123token",
			expectError:   false,
		},
		{
			name:        "Missing Authorization header",
			authHeader:  "",
			expectError: true,
		},
		{
			name:        "Authorization header with wrong prefix",
			authHeader:  "Basic abc123token",
			expectError: true,
		},
		{
			name:        "Bearer token empty",
			authHeader:  "Bearer ",
			expectError: true,
		},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest("GET", "http://example.com", nil)
		if tt.authHeader != "" {
			req.Header.Set("Authorization", tt.authHeader)
		}
		token, err := GetBearerToken(req)
		if tt.expectError {
			if err == nil {
				t.Errorf("%s: expected error but got none", tt.name)
			}
		} else {
			if err != nil {
				t.Errorf("%s: unexpected error: %v", tt.name, err)
			}
			if token != tt.expectedToken {
				t.Errorf("%s: expected token %q but got %q", tt.name, tt.expectedToken, token)
			}
		}
	}
}
