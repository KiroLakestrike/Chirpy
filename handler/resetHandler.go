package handler

import (
	"net/http"
)

func (cfg *ApiConfig) ResetHandler(w http.ResponseWriter, r *http.Request) {
	// Check if platform is dev
	if cfg.Platform != "dev" {
		RespondWithError(w, http.StatusForbidden, "Forbidden", nil)
		return
	}

	// Reset the fileserver hits
	cfg.FileserverHits.Store(0)

	// Delete all users from the database
	err := cfg.DB.DeleteAllUsers(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to delete users", err)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0 and all users deleted"))
}
