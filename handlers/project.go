package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"portfolio-v2/database"
	"portfolio-v2/models"
	"portfolio-v2/templates"
)

const projectsPerPage = 3

// ProjectsAPIHandler handles paginated project requests for HTMX
func ProjectsAPIHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 1 {
			page = 1
		}

		projects, err := database.GetProjects(db, page, projectsPerPage)
		if err != nil {
			log.Printf("Error fetching projects: %v", err)
			http.Error(w, "Error loading projects", http.StatusInternalServerError)
			return
		}

		totalProjects, err := database.CountProjects(db)
		if err != nil {
			log.Printf("Error counting projects: %v", err)
			totalProjects = 0
		}

		hasMore := (page * projectsPerPage) < totalProjects
		nextPage := page + 1

		component := templates.ProjectList(projects, hasMore, nextPage)
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Printf("Template rendering error: %v", err)
			return
		}
	}
}

// GetInitialProjects fetches the first page of projects for the home page
func GetInitialProjects(db *sql.DB) (projects []models.Project, hasMore bool, nextPage int) {
	projectList, err := database.GetProjects(db, 1, projectsPerPage)
	if err != nil {
		log.Printf("Error fetching initial projects: %v", err)
		return []models.Project{}, false, 1
	}

	totalProjects, err := database.CountProjects(db)
	if err != nil {
		log.Printf("Error counting projects: %v", err)
		totalProjects = 0
	}

	hasMore = projectsPerPage < totalProjects
	nextPage = 2

	return projectList, hasMore, nextPage
}
