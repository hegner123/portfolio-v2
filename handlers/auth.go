package handlers

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"portfolio-v2/ratelimit"
	"portfolio-v2/session"
	"portfolio-v2/templates"
)

// LoginPageHandler displays the login form
func LoginPageHandler(sessionStore *session.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Check if already logged in
		cookie, err := r.Cookie("session_id")
		if err == nil {
			if _, exists := sessionStore.Get(cookie.Value); exists {
				// Already logged in, redirect to admin
				http.Redirect(w, r, "/admin", http.StatusSeeOther)
				return
			}
		}

		// Generate CSRF token for the form
		csrfToken, err := session.GenerateCSRFToken()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Failed to generate CSRF token: %v", err)
			return
		}

		// Store CSRF token in a temporary cookie for validation
		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    csrfToken,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   600, // 10 minutes
		})

		// Get redirect target from query param
		redirectTo := r.URL.Query().Get("redirect")
		if redirectTo == "" {
			redirectTo = "/admin"
		}

		// Render login template
		component := templates.Login(csrfToken, "", redirectTo)
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Printf("Template rendering error: %v", err)
		}
	}
}

// LoginHandler handles login form submission
func LoginHandler(sessionStore *session.Store, rateLimiter *ratelimit.Limiter, hashedPassword []byte, adminUsername string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Get client IP
		ip := getClientIP(r)

		// Check rate limit
		if !rateLimiter.Allow(ip) {
			log.Printf("Rate limit exceeded for IP: %s", ip)
			component := templates.Login("", "Too many failed attempts. Please try again later.", r.FormValue("redirect"))
			component.Render(r.Context(), w)
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		// Parse form
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}

		// Validate CSRF token
		formCSRF := r.FormValue("csrf_token")
		cookieCSRF, err := r.Cookie("csrf_token")
		if err != nil || formCSRF != cookieCSRF.Value {
			log.Printf("CSRF token validation failed from IP: %s", ip)
			http.Error(w, "Invalid CSRF token", http.StatusForbidden)
			return
		}

		// Get credentials
		username := strings.TrimSpace(r.FormValue("username"))
		password := r.FormValue("password")
		redirectTo := r.FormValue("redirect")

		// Validate redirect to prevent open redirect attacks
		if redirectTo == "" || !strings.HasPrefix(redirectTo, "/admin") {
			redirectTo = "/admin"
		}

		// Validate credentials
		if username != adminUsername {
			rateLimiter.Record(ip)
			log.Printf("Failed login attempt from IP: %s (invalid username)", ip)
			component := templates.Login(formCSRF, "Invalid username or password", redirectTo)
			component.Render(r.Context(), w)
			return
		}

		if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
			rateLimiter.Record(ip)
			log.Printf("Failed login attempt from IP: %s (invalid password)", ip)
			component := templates.Login(formCSRF, "Invalid username or password", redirectTo)
			component.Render(r.Context(), w)
			return
		}

		// Successful login
		sess, err := sessionStore.Create(username)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Failed to create session: %v", err)
			return
		}

		// Reset rate limiter for this IP
		rateLimiter.Reset(ip)

		// Set session cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sess.ID,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			// No MaxAge set - session cookie (persists until browser close or logout)
		})

		// Clear CSRF cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			MaxAge:   -1,
		})

		log.Printf("Successful login for user: %s from IP: %s", username, ip)

		// Redirect to intended page
		http.Redirect(w, r, redirectTo, http.StatusSeeOther)
	}
}

// LogoutHandler handles logout requests
func LogoutHandler(sessionStore *session.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get session cookie
		cookie, err := r.Cookie("session_id")
		if err == nil {
			// Delete session from store
			sessionStore.Delete(cookie.Value)
			log.Printf("User logged out, session: %s", cookie.Value)
		}

		// Clear session cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			MaxAge:   -1,
		})

		// Redirect to home page
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// getClientIP extracts the client IP address from the request
func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header first (for proxies/load balancers)
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// Take the first IP if multiple are present
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}

	// Check X-Real-IP header
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// Fall back to RemoteAddr
	ip := r.RemoteAddr
	// Remove port if present
	if idx := strings.LastIndex(ip, ":"); idx != -1 {
		ip = ip[:idx]
	}
	return ip
}
