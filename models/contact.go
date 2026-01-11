package models

import "time"

// ContactSubmission represents a contact form submission
type ContactSubmission struct {
	ID          int64
	Name        string
	Email       string
	Message     string
	SubmittedAt time.Time
	IPAddress   string
	UserAgent   string
}
