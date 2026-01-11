package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"portfolio-v2/database"
	"portfolio-v2/handlers"
	"portfolio-v2/templates"
)

var db *sql.DB

func main() {
	var err error

	// Initialize database
	db, err = database.InitDB("./portfolio.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Seed database with sample blog posts
	if err := database.SeedBlogPosts(db); err != nil {
		log.Printf("Warning: Failed to seed blog posts: %v", err)
	}

	// Seed database with sample projects
	if err := database.SeedProjects(db); err != nil {
		log.Printf("Warning: Failed to seed projects: %v", err)
	}

	// Static file server
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/blog/", handlers.BlogPostViewHandler(db))
	http.HandleFunc("/project/", handlers.ProjectViewHandler(db))
	http.HandleFunc("/new-blog", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.NewBlogPageHandler(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateBlogPostHandler(db)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/api/blog/posts", handlers.BlogPostsAPIHandler(db))
	http.HandleFunc("/api/projects", handlers.ProjectsAPIHandler(db))

	// Start server
	port := "8080"
	fmt.Printf("Server starting on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	posts, hasMore, nextPage, tags := handlers.GetInitialBlogPosts(db)
	projects, projectsHasMore, projectsNextPage := handlers.GetInitialProjects(db)

	component := templates.Home(posts, hasMore, nextPage, tags, projects, projectsHasMore, projectsNextPage)
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template rendering error: %v", err)
		return
	}
}
