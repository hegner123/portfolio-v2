package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"portfolio-v2/database"
	"portfolio-v2/models"
	"portfolio-v2/templates"
)

// EditProjectPageHandler shows the edit form for a project
func EditProjectPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract ID from URL path /admin/project/{id}
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(pathParts) < 3 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		// Don't handle delete, new, or other special paths
		if pathParts[2] == "delete" || pathParts[2] == "new" {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(pathParts[2], 10, 64)
		if err != nil {
			http.Error(w, "Invalid project ID", http.StatusBadRequest)
			return
		}

		// Get project from database
		project, err := database.GetProjectByID(db, int(id))
		if err != nil {
			log.Printf("Error fetching project: %v", err)
			http.Error(w, "Error fetching project", http.StatusInternalServerError)
			return
		}

		if project == nil {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}

		// Render edit form
		component := templates.EditProjectForm(project)
		component.Render(r.Context(), w)
	}
}

// UpdateProjectHandler handles updating a project
func UpdateProjectHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract ID from URL path /admin/project/{id}
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(pathParts) < 3 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		// Don't handle delete, new, or other special paths
		if pathParts[2] == "delete" || pathParts[2] == "new" {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(pathParts[2], 10, 64)
		if err != nil {
			http.Error(w, "Invalid project ID", http.StatusBadRequest)
			return
		}

		// Parse form data
		err = r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Check for PUT method override
		if r.FormValue("_method") != "PUT" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		title := r.FormValue("title")
		description := r.FormValue("description")
		technologiesStr := r.FormValue("technologies")
		githubURL := r.FormValue("github_url")
		imageURL := r.FormValue("image_url")
		featured := r.FormValue("featured") == "true"

		// Validate required fields
		if title == "" || description == "" || technologiesStr == "" || imageURL == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		// Parse technologies (comma-separated)
		var technologies []string
		for _, tech := range strings.Split(technologiesStr, ",") {
			trimmed := strings.TrimSpace(tech)
			if trimmed != "" {
				technologies = append(technologies, trimmed)
			}
		}

		// Get existing project to preserve slug
		existingProject, err := database.GetProjectByID(db, int(id))
		if err != nil {
			log.Printf("Error fetching existing project: %v", err)
			http.Error(w, "Error fetching project", http.StatusInternalServerError)
			return
		}

		if existingProject == nil {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}

		// Create updated project
		project := &models.Project{
			ID:           id,
			Title:        title,
			Slug:         existingProject.Slug, // Preserve existing slug
			Description:  description,
			Technologies: technologies,
			GithubURL:    githubURL,
			ImageURL:     imageURL,
			Featured:     featured,
		}

		// Update project in database
		err = database.UpdateProject(db, project)
		if err != nil {
			log.Printf("Error updating project: %v", err)
			http.Error(w, fmt.Sprintf("Error updating project: %v", err), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
