package handler

import (
	"fmt"
	"net/http"
)

func (cfg *ApiConfig) ResetHandler(w http.ResponseWriter, r *http.Request) {
	if cfg.Platform != "dev" {
		RespondWithError(w, http.StatusForbidden, "Reset endpoint only available in dev environment", nil)
		return
	}

	// Delete refresh tokens first (because of foreign key constraint)
	err := cfg.DB.DeleteAllRefreshTokens(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to delete refresh tokens", err)
		return
	}

	// Then delete all users
	err = cfg.DB.DeleteAllUsers(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to delete users", err)
		return
	}

	// Reset metrics counter
	cfg.FileserverHits.Store(0)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits reset to %d and all users deleted", 0)))
}
