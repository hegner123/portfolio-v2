package models

import "time"

// BlogPost represents a blog post in the database
type BlogPost struct {
	ID          int64
	Title       string
	Slug        string
	Excerpt     string
	Content     string
	PublishedAt time.Time
	Tags        []string
	Author      string
}

// BlogPostPreview is a lightweight version for list views
type BlogPostPreview struct {
	ID          int64
	Title       string
	Slug        string
	Excerpt     string
	PublishedAt time.Time
	Tags        []string
}
