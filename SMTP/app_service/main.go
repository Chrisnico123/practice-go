package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"
)

const BASE_URL = "http://localhost:4000/send"
const PORT = ":4444"

type RequestBody struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
	Type    string   `json:"type"`
}

type ResponseBody struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	var req RequestBody
	err := json.NewDecoder(r.Body).Decode(&req);

	if err != nil{
		resp := ResponseBody{
			Success: false,
			Message: "Email not success",
			Error: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	reqBody , err := json.Marshal(req)
	if err != nil{
		resp := ResponseBody{
			Success: false,
			Message: "Email not success",
			Error: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp, err := http.Post(BASE_URL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil{
		resp := ResponseBody{
			Success: false,
			Message: "Email not success",
			Error: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	var RespBody ResponseBody
	err = json.NewDecoder(resp.Body).Decode(&RespBody)
	if err != nil{
		resp := ResponseBody{
			Success: false,
			Message: "Email not success",
			Error: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RespBody)
}

func main() {
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corsOptions.Handler(http.DefaultServeMux)

	http.HandleFunc("/send", sendEmail)

	log.Println("server running at port", PORT)
	err := http.ListenAndServe(PORT, handler)
	if err != nil {
		panic(err)
	}
}
