package main

import (
	"net/http"
)

func main() {
	// Serve the entire "web" folder
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.ListenAndServe(":8082", nil)
}