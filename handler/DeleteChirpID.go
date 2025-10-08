package handler

import (
	"Chirpy/internal/auth"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) DeleteChirpByID(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID string `json:"id"`
	}

	token, err := auth.GetBearerToken(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	validated, err := auth.ValidateJWT(token, cfg.ServerSecret)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	chirpIDStr := r.PathValue("id")
	chirpID, err := uuid.Parse(chirpIDStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	chirp, err := cfg.DB.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}

	if chirp.UserID != validated {
		RespondWithError(w, http.StatusForbidden, "Forbidden", nil)
		return
	}

	err = cfg.DB.DeleteChirpByID(r.Context(), chirpID)
	if err != nil {
		RespondWithError(w, http.StatusForbidden, "Forbidden", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
