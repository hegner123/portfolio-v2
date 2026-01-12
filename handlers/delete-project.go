package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	"portfolio-v2/database"
)

// DeleteProjectHandler handles deleting a project
func DeleteProjectHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract ID from URL path /admin/project/delete/{id}
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

		// Verify this is actually a delete path
		if len(pathParts) < 4 || pathParts[2] != "delete" {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(pathParts[3], 10, 64)
		if err != nil {
			http.Error(w, "Invalid project ID", http.StatusBadRequest)
			return
		}

		// Delete project from database
		err = database.DeleteProject(db, int(id))
		if err != nil {
			log.Printf("Error deleting project: %v", err)
			http.Error(w, "Error deleting project", http.StatusInternalServerError)
			return
		}

		log.Printf("Project %d deleted successfully", id)

		// Redirect back to admin dashboard
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
