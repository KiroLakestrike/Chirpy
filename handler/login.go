package handler

import (
	"Chirpy/internal/auth"
	"Chirpy/internal/database"
	"encoding/json"
	"net/http"
	"time"
)

func (cfg *ApiConfig) Login(w http.ResponseWriter, r *http.Request) {
	// Define structure to decode login request JSON
	type requestParameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Define structure for login response JSON including tokens
	type response struct {
		ID           string `json:"id"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		Email        string `json:"email"`
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	// Constants for token expiration times
	accessTokenExpiration := 3600          // 1 hour in seconds
	refreshTokenExpirationHours := 60 * 24 // 60 days in hours

	// Decode JSON body into request struct
	decoder := json.NewDecoder(r.Body)
	params := requestParameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request", err)
		return
	}

	// Load user by email from the database
	user, err := cfg.DB.GetUserByEMail(r.Context(), params.Email)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	// Verify password against stored hash
	match, err := auth.CheckPasswordHash(params.Password, user.HashedPasswords)
	if err != nil || !match {
		RespondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	// Calculate access token expiration duration
	accessExpireTime := time.Duration(accessTokenExpiration) * time.Second

	// Generate JWT access token with fixed expiration time
	accessToken, err := auth.MakeJWT(user.ID, cfg.ServerSecret, accessExpireTime)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to generate JWT AccessToken", err)
		return
	}

	// Generate a new secure refresh token string
	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to generate refresh token", err)
		return
	}

	// Calculate the refresh token expiration timestamp (60 days from now)
	refreshExpireTime := time.Now().Add(time.Duration(refreshTokenExpirationHours) * time.Hour)

	// Insert the refresh token record into the database
	refreshInsert := database.InsertRefreshTokenParams{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: refreshExpireTime,
	}
	err = cfg.DB.InsertRefreshToken(r.Context(), refreshInsert)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to store refresh token", err)
		return
	}

	// Respond with user data and both tokens on successful login
	RespondWithJSON(w, http.StatusOK, response{
		ID:           user.ID.String(),
		CreatedAt:    user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:    user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		Email:        user.Email,
		Token:        accessToken,
		RefreshToken: refreshToken,
	})
}
