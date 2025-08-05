package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"portfolio/data"
	"html/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles(
		filepath.Join("web", "templates", "base.html"),
		filepath.Join("web", "templates", "home.html"),
		filepath.Join("web", "templates", "header.html"),
		filepath.Join("web", "templates", "footer.html"),
	)
	if err != nil {
		fmt.Println("Template parse error:", err) 
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte("internal server error"))
		return
	}
	err = html.ExecuteTemplate(w, "base", data.GetHomeData())
	if err != nil {
		http.Error(w, "failed to render page", http.StatusInternalServerError)
	}
}
