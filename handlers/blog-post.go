package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"portfolio-v2/database"
	"portfolio-v2/templates"
)

// BlogPostViewHandler displays a single blog post by slug
func BlogPostViewHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract slug from path /blog/{slug}
		path := strings.TrimPrefix(r.URL.Path, "/blog/")
		slug := strings.TrimSpace(path)

		if slug == "" || slug == "/" {
			http.Redirect(w, r, "/#blog", http.StatusSeeOther)
			return
		}

		post, err := database.GetBlogPostBySlug(db, slug)
		if err != nil {
			log.Printf("Error fetching blog post by slug %s: %v", slug, err)
			http.Error(w, "Error loading post", http.StatusInternalServerError)
			return
		}

		if post == nil {
			component := templates.BlogPostNotFound()
			if err := component.Render(r.Context(), w); err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				log.Printf("Template rendering error: %v", err)
			}
			return
		}

		component := templates.BlogPostView(*post)
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Printf("Template rendering error: %v", err)
			return
		}
	}
}
