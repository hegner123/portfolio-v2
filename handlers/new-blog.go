package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"portfolio-v2/database"
	"portfolio-v2/templates"
)

// NewBlogPageHandler displays the new blog post form
func NewBlogPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	component := templates.NewBlog()
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template rendering error: %v", err)
		return
	}
}

// CreateBlogPostHandler handles blog post creation form submission
func CreateBlogPostHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			log.Printf("Form parse error: %v", err)
			return
		}

		title := strings.TrimSpace(r.FormValue("title"))
		excerpt := strings.TrimSpace(r.FormValue("excerpt"))
		content := strings.TrimSpace(r.FormValue("content"))
		tagsInput := strings.TrimSpace(r.FormValue("tags"))

		if title == "" || excerpt == "" || content == "" {
			http.Error(w, "Title, excerpt, and content are required", http.StatusBadRequest)
			return
		}

		var tags []string
		if tagsInput != "" {
			tagsParts := strings.Split(tagsInput, ",")
			for _, tag := range tagsParts {
				trimmedTag := strings.TrimSpace(tag)
				if trimmedTag != "" {
					tags = append(tags, trimmedTag)
				}
			}
		}

		_, err := database.CreateBlogPost(db, title, excerpt, content, tags)
		if err != nil {
			log.Printf("Error creating blog post: %v", err)
			http.Error(w, "Error creating blog post", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
