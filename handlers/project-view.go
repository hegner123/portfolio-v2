package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"portfolio-v2/database"
	"portfolio-v2/templates"
)

// ProjectViewHandler displays a single project by slug
func ProjectViewHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract slug from path /project/{slug}
		path := strings.TrimPrefix(r.URL.Path, "/project/")
		slug := strings.TrimSpace(path)

		if slug == "" || slug == "/" {
			http.Redirect(w, r, "/#projects", http.StatusSeeOther)
			return
		}

		project, err := database.GetProjectBySlug(db, slug)
		if err != nil {
			log.Printf("Error fetching project by slug %s: %v", slug, err)
			http.Error(w, "Error loading project", http.StatusInternalServerError)
			return
		}

		if project == nil {
			component := templates.ProjectNotFound()
			if err := component.Render(r.Context(), w); err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				log.Printf("Template rendering error: %v", err)
			}
			return
		}

		component := templates.ProjectView(*project)
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Printf("Template rendering error: %v", err)
			return
		}
	}
}
