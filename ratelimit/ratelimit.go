package ratelimit

import (
	"sync"
	"time"
)

// Limiter implements rate limiting using a sliding window algorithm
type Limiter struct {
	attempts map[string][]time.Time // IP -> attempt timestamps
	mu       sync.RWMutex
	max      int           // Maximum attempts allowed
	window   time.Duration // Time window for attempts
}

// NewLimiter creates a new rate limiter
func NewLimiter(max int, window time.Duration) *Limiter {
	return &Limiter{
		attempts: make(map[string][]time.Time),
		max:      max,
		window:   window,
	}
}

// Allow checks if an IP address is allowed to make a request
func (l *Limiter) Allow(ip string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-l.window)

	// Get attempts for this IP
	attempts, exists := l.attempts[ip]
	if !exists {
		return true
	}

	// Filter out attempts outside the window
	validAttempts := []time.Time{}
	for _, t := range attempts {
		if t.After(cutoff) {
			validAttempts = append(validAttempts, t)
		}
	}

	l.attempts[ip] = validAttempts

	// Check if under the limit
	return len(validAttempts) < l.max
}

// Record records a failed attempt for an IP address
func (l *Limiter) Record(ip string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	l.attempts[ip] = append(l.attempts[ip], now)
}

// Reset clears all attempts for an IP address (called on successful login)
func (l *Limiter) Reset(ip string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.attempts, ip)
}

// Cleanup removes old entries from the limiter (run periodically)
func (l *Limiter) Cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		l.mu.Lock()
		now := time.Now()
		cutoff := now.Add(-l.window)

		for ip, attempts := range l.attempts {
			validAttempts := []time.Time{}
			for _, t := range attempts {
				if t.After(cutoff) {
					validAttempts = append(validAttempts, t)
				}
			}

			if len(validAttempts) == 0 {
				delete(l.attempts, ip)
			} else {
				l.attempts[ip] = validAttempts
			}
		}
		l.mu.Unlock()
	}
}
