package store

import (
	"encoding/gob"
	"log"
	"os"
	"sync"
)


type UrlStore struct{
	urls map[string]string
	mu sync.RWMutex
	file *os.File
} 

type Record struct{
	KEY, URL string
}

var urlStore = NewUrlStore("store.gob")

// save the key and url in file
func (s *UrlStore) save(key, url string ) error{
	e := gob.NewEncoder(s.file)
	return e.Encode(Record{key, url})
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

// func NewUrlStore() *UrlStore {
// 	return &UrlStore{urls:make(map[string]string)}
// }

// Redefining the NewUrlStore to save urls in file
func NewUrlStore(filename string) *UrlStore{
	s := &UrlStore{urls:make(map[string]string)}
	f, err := os.OpenFile(filename,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err!=nil{
		log.Fatal("UrlStore: ", err)
	}
	s.file = f
	return s
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