package handlers

import (
	"encoding/json"
	"net/http"
	"tasko/internal/services"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var dto struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if dto.Email == "" || dto.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Call service layer
	userID, err := services.RegisterUser(dto.Email, dto.Password) // FIXED: Corrected call
	if err != nil {
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      userID,
		"message": "User registered successfully",
	})
}
