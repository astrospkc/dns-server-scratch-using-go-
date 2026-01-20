package main

import (
	"dnsServer/cmd/urlShortener/stores"
	"fmt"
	"net/http"
)

// Redirect, which redirects short URL requests
// Add, which handles the submission of new URLs



func main(){
	fmt.Print("this is url shortener worker")
	http.HandleFunc("/", stores.Redirect)
	http.HandleFunc("/add", stores.Add)
	http.ListenAndServe(":8080", nil)
}
