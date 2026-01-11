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

const postsPerPage = 3

// BlogPostsAPIHandler handles paginated blog post requests for HTMX
func BlogPostsAPIHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 1 {
			page = 1
		}

		tagFilter := r.URL.Query().Get("tag")

		posts, err := database.GetBlogPosts(db, page, postsPerPage, tagFilter)
		if err != nil {
			log.Printf("Error fetching blog posts: %v", err)
			http.Error(w, "Error loading posts", http.StatusInternalServerError)
			return
		}

		totalPosts, err := database.CountBlogPosts(db, tagFilter)
		if err != nil {
			log.Printf("Error counting blog posts: %v", err)
			totalPosts = 0
		}

		hasMore := (page * postsPerPage) < totalPosts
		nextPage := page + 1

		component := templates.BlogPostList(posts, hasMore, nextPage, tagFilter)
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Printf("Template rendering error: %v", err)
			return
		}
	}
}

// GetInitialBlogPosts fetches the first page of blog posts for the home page
func GetInitialBlogPosts(db *sql.DB) (posts []models.BlogPostPreview, hasMore bool, nextPage int, tags []string) {
	blogPosts, err := database.GetBlogPosts(db, 1, postsPerPage, "")
	if err != nil {
		log.Printf("Error fetching initial blog posts: %v", err)
		return []models.BlogPostPreview{}, false, 1, []string{}
	}

	totalPosts, err := database.CountBlogPosts(db, "")
	if err != nil {
		log.Printf("Error counting blog posts: %v", err)
		totalPosts = 0
	}

	allTags, err := database.GetAllTags(db)
	if err != nil {
		log.Printf("Error fetching tags: %v", err)
		allTags = []string{}
	}

	hasMore = postsPerPage < totalPosts
	nextPage = 2

	return blogPosts, hasMore, nextPage, allTags
}
