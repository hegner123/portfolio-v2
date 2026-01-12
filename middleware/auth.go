package middleware

import (
	"context"
	"log"
	"net/http"

	"portfolio-v2/session"
)

// contextKey is a custom type for context keys to avoid collisions
type contextKey string

const sessionContextKey contextKey = "session"

// SessionAuth wraps an http.HandlerFunc with session-based authentication
// If redirectUnauth is true, unauthenticated requests are redirected to login
// If false, a 401 Unauthorized response is returned
func SessionAuth(sessionStore *session.Store, redirectUnauth bool) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Extract session cookie
			cookie, err := r.Cookie("session_id")
			if err != nil {
				// No session cookie
				handleUnauthorized(w, r, redirectUnauth)
				return
			}

			// Validate session exists
			sess, exists := sessionStore.Get(cookie.Value)
			if !exists {
				// Invalid session
				handleUnauthorized(w, r, redirectUnauth)
				return
			}

			// Store session in context for handlers to access
			ctx := context.WithValue(r.Context(), sessionContextKey, sess)
			r = r.WithContext(ctx)

			// Call the next handler
			next(w, r)
		}
	}
}

// GetSession retrieves the session from the request context
func GetSession(r *http.Request) (*session.Session, bool) {
	sess, ok := r.Context().Value(sessionContextKey).(*session.Session)
	return sess, ok
}

// handleUnauthorized handles unauthorized requests
func handleUnauthorized(w http.ResponseWriter, r *http.Request, redirect bool) {
	if redirect {
		// Redirect to login with the current path as redirect target
		redirectURL := "/admin/login?redirect=" + r.URL.Path
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
		log.Printf("Unauthorized access attempt to %s from %s, redirecting to login", r.URL.Path, r.RemoteAddr)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		log.Printf("Unauthorized access attempt to %s from %s", r.URL.Path, r.RemoteAddr)
	}
}

/*
// ROLLBACK: Old BasicAuth implementation (kept for reference)
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
*/
