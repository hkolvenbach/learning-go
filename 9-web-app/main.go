package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	html := "<strong>Hello, world</strong>"
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, html)
}

func main() {
	http.HandleFunc("/", homepage)

	log.Println("Starting web server on port 8080")

	// Reachable via http://127.0.0.1:8080/
	http.ListenAndServe(":8080", nil)
}
