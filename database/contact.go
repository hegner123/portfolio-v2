package database

import (
	"database/sql"
	"fmt"
	"time"

	"portfolio-v2/models"
)

// CreateContactSubmission saves a contact form submission to the database
func CreateContactSubmission(db *sql.DB, name, email, message, ipAddress, userAgent string) (int64, error) {
	query := `
		INSERT INTO contact_submissions (name, email, message, submitted_at, ip_address, user_agent)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := db.Exec(query, name, email, message, time.Now(), ipAddress, userAgent)
	if err != nil {
		return 0, fmt.Errorf("insert contact submission: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}

	return id, nil
}

// GetContactSubmissions retrieves paginated contact submissions (for admin view - future)
func GetContactSubmissions(db *sql.DB, page, limit int) ([]models.ContactSubmission, error) {
	offset := (page - 1) * limit

	query := `
		SELECT id, name, email, message, submitted_at, ip_address, user_agent
		FROM contact_submissions
		ORDER BY submitted_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query contact submissions: %w", err)
	}
	defer rows.Close()

	var submissions []models.ContactSubmission
	for rows.Next() {
		var submission models.ContactSubmission
		var ipAddress, userAgent sql.NullString

		err := rows.Scan(
			&submission.ID,
			&submission.Name,
			&submission.Email,
			&submission.Message,
			&submission.SubmittedAt,
			&ipAddress,
			&userAgent,
		)
		if err != nil {
			return nil, fmt.Errorf("scan contact submission: %w", err)
		}

		if ipAddress.Valid {
			submission.IPAddress = ipAddress.String
		}
		if userAgent.Valid {
			submission.UserAgent = userAgent.String
		}

		submissions = append(submissions, submission)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate contact submissions: %w", err)
	}

	return submissions, nil
}

// CountContactSubmissions returns the total number of submissions
func CountContactSubmissions(db *sql.DB) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM contact_submissions`

	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count contact submissions: %w", err)
	}
	return count, nil
}
