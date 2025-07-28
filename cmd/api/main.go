package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sibelephant/prisma/db"
	"github.com/sibelephant/internal/config"
	"github.com/sibelephant/internal/database"
)

type application struct {
	port      int
	jwtSecret string
	models    *db.PrismaClient
}

func main() {
	// Load environment variables
	cfg := config.LoadConfig()

	// Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer database.Disconnect()

	// Create application instance
	app := &application{
		port:      cfg.Port,
		jwtSecret: cfg.JWTSecret,
		models:    database.Client, // Use Prisma client directly
	}

	log.Printf("Starting server on port %d", app.port)

	// Start the HTTP server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.port),
		Handler: app.routes(),
	}

	log.Printf("Server listening on http://localhost:%d", app.port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
