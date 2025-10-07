package handler

import (
	"Chirpy/internal/auth"
	"net/http"
)

func (cfg *ApiConfig) RevokeToken(w http.ResponseWriter, r *http.Request) {
	// Extract Refresh token from request
	bearerToken, err := auth.GetBearerToken(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Missing or invalid Authorization header", nil)
		return
	}

	err = cfg.DB.RevokeRefreshToken(r.Context(), bearerToken)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to revoke token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
