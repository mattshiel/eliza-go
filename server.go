
package main

import (
	"fmt"
	"net/http"
	"./eliza"
)

func inputhandler(w http.ResponseWriter, r *http.Request) {
	question := r.URL.Query().Get("value") // Extract question from GET request

	// Return Eliza's response's so long as the user doesn't give a quit statement
		answer := eliza.ReplyTo(question)
		
		// Return Eliza's answer
		fmt.Fprintf(w,"%s",answer)
}

func main() {
	// Adapted from https://ianmcloughlin.github.io
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", inputhandler)
	http.ListenAndServe(":8086", nil)
}