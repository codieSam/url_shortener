package services

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
)

type URLService struct {
	mu   sync.RWMutex
	urls map[string]string
}

func NewURLService() *URLService {
	return &URLService{
		urls: make(map[string]string),
	}
}

func (s *URLService) Create(originalURL string) (string, error) {
	for {
		randomBytes := make([]byte, 6)

		if _, err := rand.Read(randomBytes); err != nil {
			return "", err
		}
		code := base64.RawURLEncoding.EncodeToString(randomBytes)
		s.mu.Lock()

		if _, exists := s.urls[code]; !exists {
			s.urls[code] = originalURL
			s.mu.Unlock()
			return code, nil
		}
		s.mu.Unlock()
	}
}

// resolve function

func (s *URLService) Resolve(code string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	originalUrl, exists := s.urls[code]
	return originalUrl, exists
}
