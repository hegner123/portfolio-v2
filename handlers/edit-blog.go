package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"portfolio-v2/database"
	"portfolio-v2/templates"
)

// EditBlogPageHandler shows the edit form for a blog post
func EditBlogPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract ID from URL path /admin/blog/{id}
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

		id, err := strconv.Atoi(pathParts[2])
		if err != nil {
			http.Error(w, "Invalid blog post ID", http.StatusBadRequest)
			return
		}

		// Get blog post from database
		post, err := database.GetBlogPostByID(db, id)
		if err != nil {
			log.Printf("Error fetching blog post: %v", err)
			http.Error(w, "Error fetching blog post", http.StatusInternalServerError)
			return
		}

		if post == nil {
			http.Error(w, "Blog post not found", http.StatusNotFound)
			return
		}

		// Render edit form
		component := templates.EditBlog(post)
		component.Render(r.Context(), w)
	}
}

// UpdateBlogHandler handles updating a blog post
func UpdateBlogHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract ID from URL path /admin/blog/{id}
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

		id, err := strconv.Atoi(pathParts[2])
		if err != nil {
			http.Error(w, "Invalid blog post ID", http.StatusBadRequest)
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
		excerpt := r.FormValue("excerpt")
		content := r.FormValue("content")
		tagsStr := r.FormValue("tags")

		// Validate required fields
		if title == "" || excerpt == "" || content == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		// Parse tags (comma-separated)
		var tags []string
		if tagsStr != "" {
			for _, tag := range strings.Split(tagsStr, ",") {
				trimmed := strings.TrimSpace(tag)
				if trimmed != "" {
					tags = append(tags, trimmed)
				}
			}
		}

		// Update blog post in database
		err = database.UpdateBlogPost(db, id, title, excerpt, content, tags)
		if err != nil {
			log.Printf("Error updating blog post: %v", err)
			http.Error(w, fmt.Sprintf("Error updating blog post: %v", err), http.StatusInternalServerError)
			return
		}

		// Render success page
		component := templates.EditBlogSuccess()
		component.Render(r.Context(), w)
	}
}
