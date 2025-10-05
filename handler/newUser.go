package handler

import (
	"encoding/json"
	"net/http"
)

func (cfg *ApiConfig) NewUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}

	type response struct {
		ID        string `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Email     string `json:"email"`
	}

	// Decode the request body
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request", err)
		return
	}

	// create user in database
	user, err := cfg.DB.CreateUser(r.Context(), params.Email)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}

	// Return the user with 201 Created status
	RespondWithJSON(w, http.StatusCreated, response{
		ID:        user.ID.String(),
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		Email:     user.Email,
	})
}
