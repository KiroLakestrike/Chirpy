package handler

import (
	"net/http"

	"github.com/google/uuid"
)

// GetChirp handles GET /api/chirps/{id}
func (cfg *ApiConfig) GetChirp(w http.ResponseWriter, r *http.Request) {
	// Parse ID aus URL
	chirpIDStr := r.PathValue("id")
	chirpID, err := uuid.Parse(chirpIDStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	// Chirp aus Datenbank laden
	chirp, err := cfg.DB.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}

	// Response
	response := ChirpResponse{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}

	RespondWithJSON(w, http.StatusOK, response)
}
