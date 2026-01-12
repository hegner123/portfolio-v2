package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"portfolio-v2/database"
	"portfolio-v2/handlers"
	"portfolio-v2/middleware"
	"portfolio-v2/ratelimit"
	"portfolio-v2/session"
	"portfolio-v2/templates"
)

var db *sql.DB

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Hash admin password on startup
	adminUser := os.Getenv("ADMIN_USERNAME")
	adminPass := os.Getenv("ADMIN_PASSWORD")
	if adminUser == "" || adminPass == "" {
		log.Fatal("ADMIN_USERNAME and ADMIN_PASSWORD must be set in .env file")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(adminPass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Initialize session store (no timeout - sessions persist until logout)
	sessionStore := session.NewStore()

	// Initialize rate limiter (5 attempts, 15 minute window)
	rateLimiter := ratelimit.NewLimiter(5, 15*time.Minute)
	go rateLimiter.Cleanup()

	// Initialize database
	db, err = database.InitDB("./portfolio.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Seed database with sample projects
	if err := database.SeedProjects(db); err != nil {
		log.Printf("Warning: Failed to seed projects: %v", err)
	}

	// Custom ServeMux for 404 handling
	mux := http.NewServeMux()

	// Static file server
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/blog/", handlers.BlogPostViewHandler(db))
	mux.HandleFunc("/project/", handlers.ProjectViewHandler(db))

	// Authentication routes
	mux.HandleFunc("/admin/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.LoginPageHandler(sessionStore)(w, r)
		} else if r.Method == http.MethodPost {
			handlers.LoginHandler(sessionStore, rateLimiter, hashedPassword, adminUser)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/admin/logout", handlers.LogoutHandler(sessionStore))

	// Admin routes - protected with session authentication
	mux.HandleFunc("/admin", middleware.SessionAuth(sessionStore, true)(handlers.AdminDashboardHandler(db)))
	mux.HandleFunc("/admin/blog/new", middleware.SessionAuth(sessionStore, true)(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.NewBlogPageHandler(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateBlogPostHandler(db)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))
	mux.HandleFunc("/admin/project/new", middleware.SessionAuth(sessionStore, true)(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.NewProjectPageHandler(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateProjectHandler(db)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))
	// Edit routes - protected with session authentication
	mux.HandleFunc("/admin/blog/delete/", middleware.SessionAuth(sessionStore, true)(handlers.DeleteBlogHandler(db)))
	mux.HandleFunc("/admin/blog/", middleware.SessionAuth(sessionStore, true)(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.EditBlogPageHandler(db)(w, r)
		} else if r.Method == http.MethodPost {
			handlers.UpdateBlogHandler(db)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))
	mux.HandleFunc("/admin/project/delete/", middleware.SessionAuth(sessionStore, true)(handlers.DeleteProjectHandler(db)))
	mux.HandleFunc("/admin/project/", middleware.SessionAuth(sessionStore, true)(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.EditProjectPageHandler(db)(w, r)
		} else if r.Method == http.MethodPost {
			handlers.UpdateProjectHandler(db)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))
	mux.HandleFunc("/api/blog/posts", handlers.BlogPostsAPIHandler(db))
	mux.HandleFunc("/api/projects", handlers.ProjectsAPIHandler(db))

	// Wrap mux with 404 handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a custom ResponseWriter to capture the status code
		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		mux.ServeHTTP(rw, r)

		// If the status is 404 and nothing was written, show custom 404 page
		if rw.status == http.StatusNotFound && !rw.written {
			handlers.NotFoundHandler(w, r)
		}
	})

	// Start server
	port := "8080"
	fmt.Printf("Server starting on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	status       int
	headerWritten bool
	written      bool
}

func (rw *responseWriter) WriteHeader(status int) {
	if !rw.headerWritten {
		rw.status = status
		rw.headerWritten = true
		rw.ResponseWriter.WriteHeader(status)
	}
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	// If header wasn't written yet, write it now with current status
	if !rw.headerWritten {
		rw.WriteHeader(rw.status)
	}
	rw.written = true
	return rw.ResponseWriter.Write(b)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Only handle exact "/" path, not catch-all
	if r.URL.Path != "/" {
		// Set 404 status but don't write anything - let the wrapper handle it
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	posts, hasMore, nextPage, tags := handlers.GetInitialBlogPosts(db)
	projects, projectsHasMore, projectsNextPage := handlers.GetInitialProjects(db)

	component := templates.Home(posts, hasMore, nextPage, tags, projects, projectsHasMore, projectsNextPage)
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template rendering error: %v", err)
		return
	}
}
