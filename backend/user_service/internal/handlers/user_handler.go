package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/samoei-ftw/specno/backend/user_service/internal/services"

	auth "github.com/samoei-ftw/specno/backend/user_service/pkg"

	"golang.org/x/crypto/bcrypt"
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

	userID, err := services.RegisterUser(dto.Email, dto.Password)
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var dto struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserByEmail(dto.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	token, err := auth.GenerateToken(int(user.ID)) // TODO: add role as arg
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return token
	response := struct {
		Token  string `json:"token"`
		Role   string `json:"role"`
		UserId int    `json:"user_id"`
	}{
		Token:  token,
		Role:   user.Role,
		UserId: int(user.ID),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func FetchUserHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userIDStr, exists := vars["arg"]
	if !exists {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Convert userID to int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetch user from the service layer
	user, err := services.GetUserByID(userID)
	if err != nil {
		log.Printf("Error fetching user %d: %v", userID, err)
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Respond with user data in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
