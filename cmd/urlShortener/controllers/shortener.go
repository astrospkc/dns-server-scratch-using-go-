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


func (s *UrlStore) Set(key, url string) bool{
	s.mu.Lock()
	_,present:=s.urls[key]
	
	if present{
		s.mu.Unlock()
		return false
	}
	s.urls[key] = url
	s.mu.Unlock()
	return true

}