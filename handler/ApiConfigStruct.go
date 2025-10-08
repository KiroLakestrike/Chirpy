package handler

import (
	"Chirpy/internal/database"
	"sync/atomic"
)

type ApiConfig struct {
	FileserverHits atomic.Int32
	DB             *database.Queries
	Platform       string
	ServerSecret   string
	PolkaKey       string
}
