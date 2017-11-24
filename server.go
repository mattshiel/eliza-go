
package main

import (
	"fmt"
	"net/http"
	"./eliza"
)

func inputhandler(w http.ResponseWriter, r *http.Request) {
	question := r.URL.Query().Get("value") // Extract question from GET request

	// Create Eliza's response
	answer := eliza.GenerateAnswer(question)

	// Return the response
	fmt.Fprintf(w,"%s",answer)
}

func main() {

	// Adapted from https://ianmcloughlin.github.io
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", inputhandler)
	http.ListenAndServe(":8083", nil)
}