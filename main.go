package main

import (
	"Chirpy/handler"
	"Chirpy/internal/database"
	"Chirpy/server"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	serverSecret := os.Getenv("SERVER_SECRET")

	const filepathRoot = "."
	const port = 8080

	// Open Database Connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// Create API config
	cfg := &handler.ApiConfig{
		DB:           dbQueries,
		Platform:     platform,
		ServerSecret: serverSecret,
	}

	// Setup routes with configuration
	routes := server.SetupRoutes(server.RouteConfig{
		APIConfig:      cfg,
		FileRoot:       filepathRoot,
		HelloHandler:   handler.Hello,
		AboutHandler:   handler.About,
		HealthzHandler: handler.Healthz,
	})

	// Create and start server
	srv := server.NewServer(port, routes)

	log.Printf("Serving files from %s on port %d\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
