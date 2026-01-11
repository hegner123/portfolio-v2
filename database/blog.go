package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"portfolio-v2/models"
)

// GetBlogPosts retrieves paginated blog posts, optionally filtered by tag
func GetBlogPosts(db *sql.DB, page, limit int, tagFilter string) ([]models.BlogPostPreview, error) {
	offset := (page - 1) * limit

	query := `
		SELECT id, title, slug, excerpt, published_at, tags
		FROM blog_posts
		WHERE published_at <= ?
	`

	args := []any{time.Now()}

	if tagFilter != "" {
		query += ` AND tags LIKE ?`
		args = append(args, "%\""+tagFilter+"\"%")
	}

	query += `
		ORDER BY published_at DESC
		LIMIT ? OFFSET ?
	`

	args = append(args, limit, offset)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("query blog posts: %w", err)
	}
	defer rows.Close()

	var posts []models.BlogPostPreview
	for rows.Next() {
		var post models.BlogPostPreview
		var tagsJSON string

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Slug,
			&post.Excerpt,
			&post.PublishedAt,
			&tagsJSON,
		)
		if err != nil {
			return nil, fmt.Errorf("scan blog post: %w", err)
		}

		if err := json.Unmarshal([]byte(tagsJSON), &post.Tags); err != nil {
			post.Tags = []string{}
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate blog posts: %w", err)
	}

	return posts, nil
}

// GetBlogPostBySlug retrieves a single blog post by slug
func GetBlogPostBySlug(db *sql.DB, slug string) (*models.BlogPost, error) {
	query := `
		SELECT id, title, slug, excerpt, content, published_at, tags, author
		FROM blog_posts
		WHERE slug = ? AND published_at <= ?
	`

	var post models.BlogPost
	var tagsJSON string

	err := db.QueryRow(query, slug, time.Now()).Scan(
		&post.ID,
		&post.Title,
		&post.Slug,
		&post.Excerpt,
		&post.Content,
		&post.PublishedAt,
		&tagsJSON,
		&post.Author,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("query blog post: %w", err)
	}

	if err := json.Unmarshal([]byte(tagsJSON), &post.Tags); err != nil {
		post.Tags = []string{}
	}

	return &post, nil
}

// CountBlogPosts returns total number of published posts, optionally filtered by tag
func CountBlogPosts(db *sql.DB, tagFilter string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM blog_posts WHERE published_at <= ?`
	args := []any{time.Now()}

	if tagFilter != "" {
		query += ` AND tags LIKE ?`
		args = append(args, "%\""+tagFilter+"\"%")
	}

	err := db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count blog posts: %w", err)
	}
	return count, nil
}

// GetAllTags retrieves all unique tags from published blog posts
func GetAllTags(db *sql.DB) ([]string, error) {
	query := `
		SELECT DISTINCT tags
		FROM blog_posts
		WHERE published_at <= ?
		ORDER BY published_at DESC
	`

	rows, err := db.Query(query, time.Now())
	if err != nil {
		return nil, fmt.Errorf("query tags: %w", err)
	}
	defer rows.Close()

	tagMap := make(map[string]bool)
	for rows.Next() {
		var tagsJSON string
		if err := rows.Scan(&tagsJSON); err != nil {
			return nil, fmt.Errorf("scan tags: %w", err)
		}

		var tags []string
		if err := json.Unmarshal([]byte(tagsJSON), &tags); err == nil {
			for _, tag := range tags {
				tagMap[tag] = true
			}
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tags: %w", err)
	}

	uniqueTags := make([]string, 0, len(tagMap))
	for tag := range tagMap {
		uniqueTags = append(uniqueTags, tag)
	}

	return uniqueTags, nil
}

// CreateBlogPost inserts a new blog post into the database
func CreateBlogPost(db *sql.DB, title, excerpt, content string, tags []string) (string, error) {
	slug := generateSlug(title)
	tagsJSON, err := json.Marshal(tags)
	if err != nil {
		return "", fmt.Errorf("marshal tags: %w", err)
	}

	query := `
		INSERT INTO blog_posts (title, slug, excerpt, content, published_at, tags, author)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err = db.Exec(query, title, slug, excerpt, content, time.Now(), string(tagsJSON), "Michael")
	if err != nil {
		return "", fmt.Errorf("insert blog post: %w", err)
	}

	return slug, nil
}

// generateSlug creates a URL-friendly slug from a title
func generateSlug(title string) string {
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
