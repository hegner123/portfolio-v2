package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"time"

	"portfolio-v2/database"
	"portfolio-v2/models"
	"portfolio-v2/templates"
)

// NewProjectPageHandler displays the new project form
func NewProjectPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	component := templates.NewProjectForm()
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template rendering error: %v", err)
		return
	}
}

// CreateProjectHandler handles project creation form submission
func CreateProjectHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			component := templates.NewProjectError("Invalid form data")
			component.Render(r.Context(), w)
			log.Printf("Form parse error: %v", err)
			return
		}

		title := strings.TrimSpace(r.FormValue("title"))
		slug := strings.TrimSpace(r.FormValue("slug"))
		description := strings.TrimSpace(r.FormValue("description"))
		technologiesInput := strings.TrimSpace(r.FormValue("technologies"))
		githubURL := strings.TrimSpace(r.FormValue("github_url"))
		imageURL := strings.TrimSpace(r.FormValue("image_url"))
		featured := r.FormValue("featured") == "true"

		if title == "" || slug == "" || description == "" || imageURL == "" {
			component := templates.NewProjectError("Title, slug, description, and image URL are required")
			component.Render(r.Context(), w)
			return
		}

		var technologies []string
		if technologiesInput != "" {
			techParts := strings.Split(technologiesInput, ",")
			for _, tech := range techParts {
				trimmedTech := strings.TrimSpace(tech)
				if trimmedTech != "" {
					technologies = append(technologies, trimmedTech)
				}
			}
		}

		project := &models.Project{
			Title:        title,
			Slug:         slug,
			Description:  description,
			Technologies: technologies,
			GithubURL:    githubURL,
			ImageURL:     imageURL,
			Featured:     featured,
			CreatedAt:    time.Now(),
		}

		if err := database.CreateProject(db, project); err != nil {
			log.Printf("Error creating project: %v", err)
			component := templates.NewProjectError("Error creating project. The slug may already exist.")
			component.Render(r.Context(), w)
			return
		}

		component := templates.NewProjectSuccess(title)
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Printf("Template rendering error: %v", err)
			return
		}
	}
}
