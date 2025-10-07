package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT_Valid(t *testing.T) {
	// Arrange
	userID := uuid.New()         //Crestes a new UUID
	tokenSecret := "secret"      //Sets a Secret for the AccessToken
	expiresIn := 2 * time.Minute // Expires in 2 Minutes.

	// Act
	token, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	gotID, err := ValidateJWT(token, tokenSecret)
	if err != nil {
		t.Fatalf("ValidateJWT failed: %v", err)
	}

	// Assert
	if gotID != userID {
		t.Errorf("expected userID %v, got %v", userID, gotID)
	}
}

func TestMakeAndValidateJWT_Expired(t *testing.T) {
	// Arrange
	userID := uuid.New()
	tokenSecret := "secret"
	expiresIn := -1 * time.Minute

	// Act
	token, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	_, err = ValidateJWT(token, tokenSecret)

	// Assert
	if err == nil {
		t.Errorf("expected error, got nil")
	}

}
