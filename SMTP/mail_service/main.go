package main

import (
	"encoding/json"
	"log"
	"nbid/mail_service/config"
	"net/http"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gopkg.in/gomail.v2"
)

const (
	CONFIG_SMTP_HOST = "HOST"
	CONFIG_SMTP_PORT = "PORT"
	CONFIG_SENDER_NAME = "Chrisnico"
	CONFIG_AUTH_EMAIL = "EMAIL"
	CONFIG_AUTH_PASSWORD = "PASSWORD"
)

type RequestBody struct {
	From    string   `json:"from"`
	Subject string   `json:"subject"`
	To      []string `json:"to"`
	Type    string   `json:"type"`
	Message string   `json:"message"`
}

type ResponseBody struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func sendMail(w http.ResponseWriter , r *http.Request){
	var req RequestBody
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		resp := ResponseBody{
			Success: false,
			Message: "Email not success",
			Error: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	err = sendGoMail(req.To, req.Subject, req.Message,req.From)

	if err != nil {
		resp := ResponseBody{
			Success: false,
			Message: "Email not success",
			Error: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := ResponseBody{
		Success: true,
		Message: "Email sent successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func sendGoMail(to []string, subject, message, from string) (err error) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From" , from)
	mailer.SetHeader("To", to...)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	smtpPortStr := config.GetString(CONFIG_SMTP_PORT)

	smtpPortInt, _:= strconv.Atoi(smtpPortStr)

	dialer := gomail.NewDialer(
		config.GetString(CONFIG_SMTP_HOST),
		smtpPortInt,
		config.GetString(CONFIG_AUTH_EMAIL),
		config.GetString(CONFIG_AUTH_PASSWORD),
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	handler := corsOptions.Handler(http.DefaultServeMux)

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		sendMail(w, r)
	})

	log.Println("Running...")
	err = http.ListenAndServe(":4000", handler)
	if err != nil {
		panic(err)
	}
}