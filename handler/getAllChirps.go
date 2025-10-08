package handler

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
)

// GetAllChirps handles GET /api/chirps
func (cfg *ApiConfig) GetAllChirps(w http.ResponseWriter, r *http.Request) {
	// Get the author_id query parameter
	authorIDStr := r.URL.Query().Get("author_id")
	// Get the sort query parameter (default to "asc")
	sortOrder := r.URL.Query().Get("sort")
	if sortOrder == "" {
		sortOrder = "asc"
	}

	var chirps []ChirpResponse

	if authorIDStr == "" {
		// No author_id provided - get ALL chirps
		dbChirps, err := cfg.DB.GetAllChirps(r.Context())
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Failed to fetch chirps", err)
			return
		}

		// Convert to response format
		chirps = make([]ChirpResponse, 0, len(dbChirps))
		for _, chirp := range dbChirps {
			chirps = append(chirps, ChirpResponse{
				ID:        chirp.ID,
				CreatedAt: chirp.CreatedAt,
				UpdatedAt: chirp.UpdatedAt,
				Body:      chirp.Body,
				UserID:    chirp.UserID,
			})
		}
	} else {
		// author_id provided - parse and filter by author
		authorID, parseErr := uuid.Parse(authorIDStr)
		if parseErr != nil {
			RespondWithError(w, http.StatusBadRequest, "Invalid author_id format", parseErr)
			return
		}

		// Get chirps for specific author
		dbChirps, err := cfg.DB.GetChirpsByAuthor(r.Context(), authorID)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Failed to fetch chirps", err)
			return
		}

		// Convert to response format
		chirps = make([]ChirpResponse, 0, len(dbChirps))
		for _, chirp := range dbChirps {
			chirps = append(chirps, ChirpResponse{
				ID:        chirp.ID,
				CreatedAt: chirp.CreatedAt,
				UpdatedAt: chirp.UpdatedAt,
				Body:      chirp.Body,
				UserID:    chirp.UserID,
			})
		}
	}

	// Sort chirps based on sort parameter
	sort.Slice(chirps, func(i, j int) bool {
		if sortOrder == "desc" {
			return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
		}
		// Default to ascending
		return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
	})

	// Return JSON response
	RespondWithJSON(w, http.StatusOK, chirps)
}
