package handler

import (
	"Chirpy/internal/auth"
	"Chirpy/internal/database"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ChirpRequest struct {
	Body string `json:"body"`
}

type ChirpResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

// Chirps handles POST /api/chirps
func (cfg *ApiConfig) Chirps(w http.ResponseWriter, r *http.Request) {
	// Authenticate user with AccessToken
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

	// Decode request
	var req ChirpRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Validate chirp body (max 140 characters)
	if len(req.Body) > 140 {
		RespondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	if len(req.Body) == 0 {
		RespondWithError(w, http.StatusBadRequest, "Chirp body cannot be empty", nil)
		return
	}

	// Clean profane words
	cleanedBody := cleanProfaneWords(req.Body)

	// Create chirp in database
	chirp, err := cfg.DB.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   cleanedBody,
		UserID: validated,
	})

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to create chirp", err)
		return
	}

	// Respond with created chirp
	response := ChirpResponse{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}

	RespondWithJSON(w, http.StatusCreated, response)
}

// cleanProfaneWords replaces profane words with ****
func cleanProfaneWords(text string) string {
	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Split(text, " ")

	for i, word := range words {
		lowerWord := strings.ToLower(word)
		for _, profane := range profaneWords {
			if lowerWord == profane {
				words[i] = "****"
				break
			}
		}
	}

	return strings.Join(words, " ")
}
