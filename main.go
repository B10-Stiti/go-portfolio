package main

import (
	"fmt"
	"net/http"
	"portfolio/handlers"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/", handlers.HomeHandler)

	fs := http.FileServer(http.Dir("web/static"))
	server.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(":5000", server)
	if err != nil {
		fmt.Printf("Failed to start server")
	}
}
