// Author: Matthew Shiel
package main

import (
	"fmt"
	"net/http"
	"./eliza"
)

func inputhandler(w http.ResponseWriter, r *http.Request) {
	// Extract question from GET request
	question := r.URL.Query().Get("value") 

	// Return Eliza's response's so long as the user doesn't give a quit statement
	answer := eliza.ReplyTo(question)
		
	// Return Eliza's answer
	fmt.Fprintf(w,"%s",answer)
}

func main() {
	// Adapted from https://ianmcloughlin.github.io
	// Serves the web file
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	// Handles the user input and return of Eliza's answers
	http.HandleFunc("/user-input", inputhandler)
	http.ListenAndServe(":8092", nil)
}