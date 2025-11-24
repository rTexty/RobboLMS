package main

import (
	"lms-backend/internal/api"
	"lms-backend/internal/store"
	"log"
	"net/http"
	"os"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/lms?sslmode=disable"
	}

	store, err := store.New(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	handler := api.NewHandler(store)
	handler.RegisterRoutes(mux)

	// CORS Middleware
	corsMux := enableCORS(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, corsMux); err != nil {
		log.Fatal(err)
	}
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
