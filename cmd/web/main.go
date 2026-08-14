package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PawBer/quizforge/internal/web"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := web.Application{}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, app.SetupRoutes()); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
