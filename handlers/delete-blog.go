package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	"portfolio-v2/database"
)

// DeleteBlogHandler handles deleting a blog post
func DeleteBlogHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract ID from URL path /admin/blog/delete/{id}
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

		// Verify this is actually a delete path
		if len(pathParts) < 4 || pathParts[2] != "delete" {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(pathParts[3])
		if err != nil {
			http.Error(w, "Invalid blog post ID", http.StatusBadRequest)
			return
		}

		// Delete blog post from database
		err = database.DeleteBlogPost(db, id)
		if err != nil {
			log.Printf("Error deleting blog post: %v", err)
			http.Error(w, "Error deleting blog post", http.StatusInternalServerError)
			return
		}

		log.Printf("Blog post %d deleted successfully", id)

		// Redirect back to admin dashboard
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
