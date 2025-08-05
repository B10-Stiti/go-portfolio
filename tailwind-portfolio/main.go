package main

import (
	"log"
	"net/http"
	"portfolio/handlers"
	"github.com/joho/godotenv"
)

func init() {
    _ = godotenv.Load()
}

func main() {
	server := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/static"))
	server.Handle("/static/", http.StripPrefix("/static/", fs))
	server.HandleFunc("/", handlers.HomeHandler)
	server.HandleFunc("/about", handlers.AboutHandler)
	server.HandleFunc("/contact", handlers.ContactHandler)
	server.HandleFunc("/api/contact", handlers.FormHandler)
	server.HandleFunc("/thankyou", handlers.ThankyouHandler)
	server.HandleFunc("/projects", handlers.ProjectHandler)

	log.Println("🚀 Server starting on http://localhost:5000")
	err := http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
