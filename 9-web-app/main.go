package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", homepage)

	log.Println("Starting web server on port 8080")

	http.ListenAndServe(":8080", nil)

}
