package controllers

import (
	"backend/models"
	"encoding/json"

	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "Application/json")
	var registerData models.Register
	err := json.NewDecoder(r.Body).Decode(&registerData)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user logged in successfully"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginData models.Login
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "Invalid Login details", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user logged in successfully"))
}
