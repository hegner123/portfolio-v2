package models

import "time"

// Project represents a project in the database
type Project struct {
	ID          int64
	Title       string
	Slug        string
	Description string
	Technologies []string
	GithubURL   string
	ImageURL    string
	Featured    bool
	CreatedAt   time.Time
}
