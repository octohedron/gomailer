package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/gorilla/mux"
)

//  The port the program will be serving on
const port = "5555"

// The email from which the emails will be sent
var fromEmail = ""

// The email password to send the emails
var emailPassword = ""

// The email to which the emails will be sent
var toEmail = ""

func init() {
	fromEmail = os.Getenv("from_email")
	emailPassword = os.Getenv("email_password")
	toEmail = os.Getenv("to_email")
}

func setCORSHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*") // careful with this, change it to suit your needs
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	return w
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	// Set headers for CORS
	setCORSHeaders(w)
	// We only allow for POST requests
	if r.Method != "POST" {
		http.Error(w, fmt.Sprintf("Forbidden request"), 403)
		return
	}
	// Set email values
	from := fromEmail
	pass := emailPassword
	to := toEmail
	// Read from the post request form data values
	clientEmail := r.FormValue("email")
	clientName := r.FormValue("name")
	subject := r.FormValue("subject")
	message := r.FormValue("message")
	// Build email
	msg := "From: " + clientName + " " + clientEmail + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		message
	// Send email
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
	// Check if there was a problem with SMTP
	if err != nil {
		http.Error(w, fmt.Sprintf("smtp error: %s", err), 400)
		return
	}
	// Write 'sent' as response
	w.Write([]byte("sent"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sendemail", sendEmail)
	log.Println("Serving email server in port " + port)
	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
