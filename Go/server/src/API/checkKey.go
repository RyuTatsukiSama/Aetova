package server_api

import (
	"net/http"
	"os"
)

// Middleware pour vérifier l'API key
func requireAPIKey(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("api_key")

		if apiKey == "" {
			http.Error(w, "API key missing", http.StatusUnauthorized)
			return
		}

		validKey := os.Getenv("VALID_KEY")
		if validKey != apiKey {
			http.Error(w, "Invalid API key", http.StatusForbidden)
			return
		}

		// Si l'API key est valide, passe à la suite
		next(w, r)
	}
}
