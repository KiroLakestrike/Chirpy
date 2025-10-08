package handler

import (
	"Chirpy/internal/auth"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) ApplyChirpyRed(w http.ResponseWriter, r *http.Request) {
	// Check the Header for a APi Key
	token, err := auth.GetAPIKey(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Invalid API Key", err)
		return
	}

	// if the API Key doesnt match the POLKA_API_KEY return 401
	if token != cfg.PolkaKey {
		RespondWithError(w, http.StatusUnauthorized, "Invalid API Key", nil)
		return
	}

	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		}
	}

	// Decode the request body
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request", err)
		return
	}

	// Check if the event is "user.upgraded" if not return 204
	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Check if the user exists
	_, err = cfg.DB.GetUserByID(r.Context(), params.Data.UserID)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "User not found", err)
		return
	}

	// Update the user in the database
	_, err = cfg.DB.ApplyChirpyRed(r.Context(), params.Data.UserID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't update user", err)
		return
	}

	// return 204 and empty body if sucessful,
	w.WriteHeader(http.StatusNoContent)
}
