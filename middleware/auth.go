package middleware

import (
	"log"
	"net/http"
	"os"

	"portfolio-v2/templates"
)

// BasicAuth wraps an http.HandlerFunc with HTTP Basic Authentication
// Credentials are read from ADMIN_USERNAME and ADMIN_PASSWORD environment variables
func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		adminUser := os.Getenv("ADMIN_USERNAME")
		adminPass := os.Getenv("ADMIN_PASSWORD")

		// If no credentials are configured, show setup instructions
		if adminUser == "" || adminPass == "" {
			log.Println("WARNING: ADMIN_USERNAME or ADMIN_PASSWORD not set")
			w.WriteHeader(http.StatusServiceUnavailable)
			component := templates.AdminSetupRequired()
			component.Render(r.Context(), w)
			return
		}

		// Validate credentials
		if !ok || username != adminUser || password != adminPass {
			w.Header().Set("WWW-Authenticate", `Basic realm="Admin Area"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Printf("Failed authentication attempt from %s", r.RemoteAddr)
			return
		}

		// Log successful authentication
		log.Printf("Admin authenticated: %s from %s", username, r.RemoteAddr)

		// Call the next handler
		next(w, r)
	}
}
