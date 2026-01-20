package main

import (
	"dnsServer/cmd/urlShortener/store"
	"fmt"
	"net/http"
)

// Redirect, which redirects short URL requests
// Add, which handles the submission of new URLs



func main(){
	fmt.Print("this is url shortener worker")
	http.HandleFunc("/", store.Redirect)
	http.HandleFunc("/add", store.Add)
	http.ListenAndServe(":8080", nil)
}
