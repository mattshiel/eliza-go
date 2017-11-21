
package main

import (
	"fmt"
	"net/http"
)

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wait, %s!", r.URL.Query().Get("value")) //.Path[1:])
}

func main() {

	// Adapted https://ianmcloughlin.github.io
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8989", nil)
}