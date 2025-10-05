package handler

import (
	"net/http"
)

// GetAllChirps handles GET /api/chirps
func (cfg *ApiConfig) GetAllChirps(w http.ResponseWriter, r *http.Request) {
	// Alle Chirps aus der Datenbank laden
	chirps, err := cfg.DB.GetAllChirps(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to fetch chirps", err)
		return
	}

	// Konvertiere zu Response-Format
	response := make([]ChirpResponse, 0, len(chirps))
	for _, chirp := range chirps {
		response = append(response, ChirpResponse{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}

	// JSON zurückgeben
	RespondWithJSON(w, http.StatusOK, response)
}
