package main

import (
	"log"
	"net/http"
	"portfolio/handlers"
)

func main() {
	server := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/static"))
	server.Handle("/static/", http.StripPrefix("/static/", fs))
	server.HandleFunc("/", handlers.HomeHandler)
	server.HandleFunc("/about", handlers.AboutHandler)

	log.Println("🚀 Server starting on http://localhost:5000")
	err := http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
