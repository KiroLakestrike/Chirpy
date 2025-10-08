package handler

import (
	"Chirpy/internal/auth"
	"Chirpy/internal/database"
	"encoding/json"
	"net/http"
	"time"
)

func (cfg *ApiConfig) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Define structs
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		ID          string    `json:"id"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Email       string    `json:"email"`
		IsChirpyRed bool      `json:"is_chirpy_red"`
	}

	// Authenticate user with AccessToken
	token, err := auth.GetBearerToken(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	// Validate JWT
	validated, err := auth.ValidateJWT(token, cfg.ServerSecret)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	// Decode the request body
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request", err)
		return
	}

	// Hash the password
	hashed, err := auth.HashPassword(params.Password)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
		return
	}

	// Create user parameters for database
	userParams := database.UpdateUserParams{
		Email:           params.Email,
		HashedPasswords: hashed,
		ID:              validated,
	}

	// Update user in database
	user, err := cfg.DB.UpdateUser(r.Context(), userParams)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't update user", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, response{
		ID:          user.ID.String(),
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Email:       user.Email,
		IsChirpyRed: user.IsChirpyRed,
	})

}
