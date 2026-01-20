package stores

import (
	"encoding/gob"
	"io"
	"log"
	"os"
	"sync"
)

// Instead of saving records directly to disk, we send them to a channel, which is a kind of buffer, so the sending function doesn’t have to wait for it.
const saveQueueLength = 1000
type UrlStore struct{
	urls map[string]string
	mu sync.RWMutex
	save chan Record
} 

var Urlstore *UrlStore

// save := make(chan, saveQueueLength)

type Record struct{
	KEY, URL string
}



// a new save method that writes a given key and URL to disk as a gob-encoded record:



// load the data stored in the disk and read in map
func (s *UrlStore) load(filename string) error{
	f, err:=os.Open(filename)
	if err!=nil{
		log.Println("Error opening file: ", err)
		return err
	}
	defer f.Close()

	d := gob.NewDecoder(f)
	
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
	
	// replacing this unnecessary file opening with go saveLoop
	// f, err := os.OpenFile(filename,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	s:=&UrlStore{
	urls:make(map[string]string),
	save:make(chan Record, saveQueueLength),
	}
	
	if err:=s.load(filename); err!=nil{
		log.Println("Error loading urlstore: ", err)
	}
	go s.SaveLoop(filename)
	return s
}

// Records are read from the save channel in an infinite loop and encoded to the file.

func (s *UrlStore) SaveLoop(filename string){
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err!=nil{
		log.Fatal("UrlStore: ", err)
	}
	defer f.Close()
	e:=gob.NewEncoder(f)
	for{
		r:=<-s.save // taking a arecord from the channel and encoding it
		if err:=e.Encode(r); err!=nil{
			log.Println("UrlStore", err)
		}
	}
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
			s.save<-Record{key, url}
			return key
		}
	}
	
}

