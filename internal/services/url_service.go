package services

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
)

type urlService struct {
	mu   sync.RWMutex
	urls map[string]string
}

func newURLService() *urlService {
	return &urlService{
		urls: make(map[string]string),
	}
}

func (s *urlService) Create(originalURL string) (string, error) {
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

func (s *urlService) Resolve(code string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	originalUrl, exists := s.urls[code]
	return originalUrl, exists
}
