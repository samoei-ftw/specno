package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/samoei-ftw/specno/backend/user_service/internal/services"

	auth "github.com/samoei-ftw/specno/backend/user_service/pkg"

	"strconv"

	"github.com/gorilla/mux"
	"github.com/samoei-ftw/specno/backend/common/utils"

	"github.com/samoei-ftw/specno/backend/gateways"
	"golang.org/x/crypto/bcrypt"
)
var userGateway = gateways.UserGatewayInit()
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

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL
	vars := mux.Vars(r)
	userId, exists := vars["arg"]
	if !exists {
		utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Missing user ID in request",
		})
		return
	}
	userID, err := strconv.Atoi(userId)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid user ID. Check if it is a valid guid.",
		})
		return
	}
	user, err := userGateway.GetUserByID(userID)
	if err != nil {
		log.Printf("Error fetching user %d: %v", userID, err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, utils.Response{
			Status:  "error",
			Message: "Failed to retrieve user",
		})
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, utils.Response{
		Status: "success",
		Data:   user,
	})
}