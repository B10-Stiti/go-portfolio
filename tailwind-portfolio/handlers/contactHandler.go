package handlers

import (
	"fmt"
	"github.com/resend/resend-go/v2"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"portfolio/data"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles(
		filepath.Join("web", "templates", "base.html"),
		filepath.Join("web", "templates", "contact.html"),
		filepath.Join("web", "templates", "header.html"),
		filepath.Join("web", "templates", "footer.html"),
	)
	if err != nil {
		fmt.Println("Template parse error:", err)
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte("internal server error"))
		return
	}
	err = html.ExecuteTemplate(w, "base", data.GetContactData())
	if err != nil {
		http.Error(w, "failed to render page", http.StatusInternalServerError)
		fmt.Println("render error :", err)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	userName := r.FormValue("name")
	userEmail := r.FormValue("email")
	message := r.FormValue("message")
	if userName == "" || userEmail == "" || message == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}
	// send the email
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		fmt.Println("Missing RESEND_API_KEY")
		http.Error(w, "Server misconfigured", http.StatusInternalServerError)
		return
	}
	client := resend.NewClient(apiKey)
	params := &resend.SendEmailRequest{
		From:    "badis_key@resend.dev",
		To:      []string{"badisstiti11@gmail.com"},
		Subject: "New message from your portfolio",
		Html: fmt.Sprintf(`<p><strong>Name:</strong> %s</p>
	                   <p><strong>Email:</strong> %s</p>
	                   <p><strong>Message:</strong><br>%s</p>`, userName, userEmail, message),
	}
	_, err = client.Emails.Send(params)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		fmt.Println("Email send error:", err)
		return
	}
	http.Redirect(w, r, "/thankyou", http.StatusSeeOther)

}
