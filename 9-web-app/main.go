package main

import (
	"log"
	"net/http"
	"text/template"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", homepage)

	log.Println("Starting web server on port 8080")

	// Reachable via http://127.0.0.1:8080/
	http.ListenAndServe(":8080", nil)
}
