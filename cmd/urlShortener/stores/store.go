package stores

import (
	"encoding/gob"
	"io"
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

// a new save method that writes a given key and URL to disk as a gob-encoded record:
func (s *UrlStore) save(key, url string ) error{
	e := gob.NewEncoder(s.file)
	return e.Encode(Record{key, url})
}


// load the data stored in the disk and read in map
func (s *UrlStore) load() error{
	if _, err:=s.file.Seek(0,0);err!=nil{
		return err
	}
	d := gob.NewDecoder(s.file)
	var err error 
	for err==nil{
		
		var r Record
		if err = d.Decode(&r); err==nil{
			s.Set(r.KEY,r.URL)
		}
	}
	if err==io.EOF{
		return nil
	}
	return err
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
	if err:= s.load(); err!=nil{
		log.Println("Error loading in dataStore", err)
	}
	return s
}




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
			if err:=s.save(key, url); err!=nil{
				log.Println("error saving to urlstore: ", err)
			}
			return key
		}
	}
	
}

