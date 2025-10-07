package handler

import (
	"Chirpy/internal/auth"
	"net/http"
	"time"
)

func (cfg *ApiConfig) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Extract Refresh token from request (without requiring "Bearer " prefix)
	refreshToken, err := auth.GetRefreshToken(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	// Call the generated sqlc method GetUserFromRefreshToken with the refresh token string.
	user, err := cfg.DB.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	// generate new JWT access Token
	accessExpireTime := time.Hour
	newAccessToken, err := auth.MakeJWT(user.ID, cfg.ServerSecret, accessExpireTime)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to generate access token", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{
		"token": newAccessToken,
	})
}
