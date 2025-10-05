package handler

import (
	"Chirpy/internal/auth"
	"encoding/json"
	"net/http"
)

func (cfg *ApiConfig) Login(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		ID        string `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Email     string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request", err)
		return
	}

	// Load user from Database
	user, err := cfg.DB.GetUserByEMail(r.Context(), params.Email)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	// Check if password matches
	match, err := auth.CheckPasswordHash(params.Password, user.HashedPasswords)
	if err != nil || !match {
		RespondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	// Success
	RespondWithJSON(w, http.StatusOK, response{
		ID:        user.ID.String(),
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		Email:     user.Email,
	})
}
