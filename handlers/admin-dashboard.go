package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"portfolio-v2/database"
	"portfolio-v2/templates"
)

// AdminDashboardHandler shows the admin dashboard with all content
func AdminDashboardHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Get all blog posts
		blogs, err := database.GetAllBlogPosts(db)
		if err != nil {
			log.Printf("Error fetching blog posts: %v", err)
			http.Error(w, "Error fetching blog posts", http.StatusInternalServerError)
			return
		}

		// Get all projects
		projects, err := database.GetAllProjects(db)
		if err != nil {
			log.Printf("Error fetching projects: %v", err)
			http.Error(w, "Error fetching projects", http.StatusInternalServerError)
			return
		}

		// Render dashboard
		component := templates.AdminDashboard(blogs, projects)
		component.Render(r.Context(), w)
	}
}
