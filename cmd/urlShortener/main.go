package main

import (
	"dnsServer/cmd/urlShortener/stores"
	"flag"
	"fmt"
	"net/http"
)

// Redirect, which redirects short URL requests
// Add, which handles the submission of new URLs

var (
  listenAddr = flag.String("http", ":3000", "http listen address")
  dataFile = flag.String("file", "store.gob", "data store file name")
  hostname = flag.String("host", "1dkne4jl5mmmm.educative.run", "host name and port")
)




func main(){
	flag.Parse()
	fmt.Print("this is url shortener worker")
	stores.Urlstore =stores.NewUrlStore(*dataFile)
	http.HandleFunc("/", stores.Redirect)
	http.HandleFunc("/add", stores.Add)
	http.ListenAndServe(":8080", nil)
}
