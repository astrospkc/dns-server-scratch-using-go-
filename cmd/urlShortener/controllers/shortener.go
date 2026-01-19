package controllers

import "sync"


type UrlStore struct{
	urls map[string]string
	mu sync.RWMutex
} 

func (s *UrlStore) Get(key string) string{
	s.mu.RLock()
	url:=s.urls[key]
	s.mu.Unlock()
	return url
}