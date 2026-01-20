package store

import (
	"os"
	"sync"
)


type UrlStore struct{
	urls map[string]string
	mu sync.RWMutex
	file *os.File
} 

func (s *UrlStore) Get(key string) string{
	s.mu.RLock()
	defer s.mu.RUnlock()
	url:=s.urls[key]
	return url
}


func (s *UrlStore) Set(key, url string) bool{
	s.mu.Lock()
	defer s.mu.Unlock()
	_,present:=s.urls[key]
	
	if present{
		s.mu.Unlock()
		return false
	}
	s.urls[key] = url
	
	return true

}

func NewUrlStore() *UrlStore {
	return &UrlStore{urls:make(map[string]string)}
}

var store = NewUrlStore()


// To add new short/loong url
// if s.Set("a", "http://google.com") {
//   // success
// }

// to retrieve
// if url := s.Get("a"); url!= ""{
// 	// redirent url
// }else{
// 	// key not founf
// }

func (s *UrlStore) Count() int{
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}


// make a Put method that takes a long URL, generates its short key with genKey, stores the URL under this (short URL) key with the Set method, and returns that key:
func (s *UrlStore) Put(url string)string{
	
	for{
		key:=GenKey(s.Count())
		if s.Set(key, url){
			return key
		}
	}


}