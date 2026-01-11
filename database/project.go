package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"portfolio-v2/models"
)

// GetProjects retrieves paginated projects, with featured projects first
func GetProjects(db *sql.DB, page, limit int) ([]models.Project, error) {
	offset := (page - 1) * limit

	query := `
		SELECT id, title, slug, description, technologies, github_url, image_url, featured, created_at
		FROM projects
		ORDER BY featured DESC, created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query projects: %w", err)
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		var technologiesJSON string
		var featured int

		err := rows.Scan(
			&project.ID,
			&project.Title,
			&project.Slug,
			&project.Description,
			&technologiesJSON,
			&project.GithubURL,
			&project.ImageURL,
			&featured,
			&project.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan project: %w", err)
		}

		project.Featured = featured == 1

		if err := json.Unmarshal([]byte(technologiesJSON), &project.Technologies); err != nil {
			project.Technologies = []string{}
		}

		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate projects: %w", err)
	}

	return projects, nil
}

// GetProjectBySlug retrieves a single project by slug
func GetProjectBySlug(db *sql.DB, slug string) (*models.Project, error) {
	query := `
		SELECT id, title, slug, description, technologies, github_url, image_url, featured, created_at
		FROM projects
		WHERE slug = ?
	`

	var project models.Project
	var technologiesJSON string
	var featured int

	err := db.QueryRow(query, slug).Scan(
		&project.ID,
		&project.Title,
		&project.Slug,
		&project.Description,
		&technologiesJSON,
		&project.GithubURL,
		&project.ImageURL,
		&featured,
		&project.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("query project: %w", err)
	}

	project.Featured = featured == 1

	if err := json.Unmarshal([]byte(technologiesJSON), &project.Technologies); err != nil {
		project.Technologies = []string{}
	}

	return &project, nil
}

// CountProjects returns total number of projects
func CountProjects(db *sql.DB) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM projects`

	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count projects: %w", err)
	}
	return count, nil
}

// SeedProjects adds sample projects for development
func SeedProjects(db *sql.DB) error {
	count, err := CountProjects(db)
	if err != nil {
		return err
	}

	if count > 0 {
		log.Println("Projects already exist, skipping seed")
		return nil
	}

	projects := []struct {
		title        string
		slug         string
		description  string
		technologies []string
		githubURL    string
		imageURL     string
		featured     bool
	}{
		{
			title:        "Portfolio V2",
			slug:         "portfolio-v2",
			description:  "A modern single-page portfolio website built with Go, Templ, and HTMX. Features interactive animations, a blog system with SQLite storage, and dynamic content loading without heavy JavaScript frameworks. Demonstrates server-side rendering with progressive enhancement.",
			technologies: []string{"Go", "Templ", "HTMX", "SQLite", "JavaScript"},
			githubURL:    "https://github.com/yourusername/portfolio-v2",
			imageURL:     "/static/images/projects/portfolio-v2.png",
			featured:     true,
		},
		{
			title:        "Task Management API",
			slug:         "task-management-api",
			description:  "RESTful API for task management built with Go and PostgreSQL. Features JWT authentication, role-based access control, and comprehensive test coverage. Includes user management, team collaboration, and task assignment workflows.",
			technologies: []string{"Go", "PostgreSQL", "JWT", "Docker", "REST API"},
			githubURL:    "https://github.com/yourusername/task-api",
			imageURL:     "/static/images/projects/task-api.png",
			featured:     true,
		},
		{
			title:        "Real-Time Chat Application",
			slug:         "realtime-chat-app",
			description:  "WebSocket-based chat application with Go backend and React frontend. Supports multiple chat rooms, direct messaging, typing indicators, and message history. Uses Redis for pub/sub and session management.",
			technologies: []string{"Go", "React", "WebSocket", "Redis", "PostgreSQL"},
			githubURL:    "https://github.com/yourusername/chat-app",
			imageURL:     "/static/images/projects/chat-app.png",
			featured:     true,
		},
		{
			title:        "E-Commerce Platform",
			slug:         "ecommerce-platform",
			description:  "Full-stack e-commerce solution with product catalog, shopping cart, and payment processing. Built with Node.js and Next.js, featuring Stripe integration, inventory management, and order tracking. Includes admin dashboard for product and order management.",
			technologies: []string{"Node.js", "Next.js", "PostgreSQL", "Stripe", "Tailwind CSS"},
			githubURL:    "https://github.com/yourusername/ecommerce",
			imageURL:     "/static/images/projects/ecommerce.png",
			featured:     false,
		},
		{
			title:        "Weather Dashboard",
			slug:         "weather-dashboard",
			description:  "Interactive weather dashboard displaying real-time weather data and forecasts. Integrates with OpenWeatherMap API, features location search, 7-day forecasts, and weather alerts. Built with vanilla JavaScript and modern CSS for responsive design.",
			technologies: []string{"JavaScript", "HTML", "CSS", "REST API"},
			githubURL:    "https://github.com/yourusername/weather-dashboard",
			imageURL:     "/static/images/projects/weather.png",
			featured:     false,
		},
		{
			title:        "CI/CD Pipeline Automation",
			slug:         "cicd-automation",
			description:  "Automated CI/CD pipeline configuration and tooling for containerized applications. Includes GitHub Actions workflows, Docker multi-stage builds, Kubernetes deployment manifests, and monitoring setup with Prometheus and Grafana.",
			technologies: []string{"Docker", "Kubernetes", "GitHub Actions", "Terraform", "Prometheus"},
			githubURL:    "https://github.com/yourusername/cicd-automation",
			imageURL:     "/static/images/projects/cicd.png",
			featured:     false,
		},
	}

	stmt, err := db.Prepare(`
		INSERT INTO projects (title, slug, description, technologies, github_url, image_url, featured)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, project := range projects {
		technologiesJSON, err := json.Marshal(project.technologies)
		if err != nil {
			return fmt.Errorf("marshal technologies: %w", err)
		}

		featuredInt := 0
		if project.featured {
			featuredInt = 1
		}

		_, err = stmt.Exec(
			project.title,
			project.slug,
			project.description,
			string(technologiesJSON),
			project.githubURL,
			project.imageURL,
			featuredInt,
		)
		if err != nil {
			return fmt.Errorf("insert project %s: %w", project.slug, err)
		}
	}

	log.Printf("Seeded %d projects", len(projects))
	return nil
}

// generateProjectSlug creates a URL-friendly slug from a title
func generateProjectSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove special characters except hyphens
	var result strings.Builder
	for _, char := range slug {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' {
			result.WriteRune(char)
		}
	}

	// Remove consecutive hyphens
	finalSlug := result.String()
	for strings.Contains(finalSlug, "--") {
		finalSlug = strings.ReplaceAll(finalSlug, "--", "-")
	}

	// Trim hyphens from start and end
	finalSlug = strings.Trim(finalSlug, "-")

	return finalSlug
}
