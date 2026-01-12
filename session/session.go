package session

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

// Session represents a user session
type Session struct {
	ID        string
	Username  string
	CreatedAt time.Time
	CSRFToken string
}

// Store manages sessions in memory
type Store struct {
	sessions map[string]*Session
	mu       sync.RWMutex
}

// NewStore creates a new session store
func NewStore() *Store {
	return &Store{
		sessions: make(map[string]*Session),
	}
}

// Create creates a new session for the given username
func (s *Store) Create(username string) (*Session, error) {
	sessionID, err := generateID()
	if err != nil {
		return nil, err
	}

	csrfToken, err := generateID()
	if err != nil {
		return nil, err
	}

	session := &Session{
		ID:        sessionID,
		Username:  username,
		CreatedAt: time.Now(),
		CSRFToken: csrfToken,
	}

	s.mu.Lock()
	s.sessions[sessionID] = session
	s.mu.Unlock()

	return session, nil
}

// Get retrieves a session by ID
func (s *Store) Get(sessionID string) (*Session, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, exists := s.sessions[sessionID]
	return session, exists
}

// Delete removes a session from the store
func (s *Store) Delete(sessionID string) {
	s.mu.Lock()
	delete(s.sessions, sessionID)
	s.mu.Unlock()
}

// generateID generates a cryptographically secure random ID
func generateID() (string, error) {
	b := make([]byte, 32) // 256 bits
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// GenerateCSRFToken generates a CSRF token for forms
func GenerateCSRFToken() (string, error) {
	return generateID()
}
