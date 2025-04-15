package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	utils "github.com/samoei-ftw/specno/backend/common/utils"
	dto "github.com/samoei-ftw/specno/backend/user_service/internal/models"
	"github.com/samoei-ftw/specno/backend/user_service/internal/services"

	auth "github.com/samoei-ftw/specno/backend/user_service/pkg"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, service *services.UserService) {
	var userReqisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userReqisterRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if userReqisterRequest.Email == "" || userReqisterRequest.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	userID, err := service.RegisterUser(userReqisterRequest.Email, userReqisterRequest.Password)
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

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request, service *services.UserService) {
	var dto struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := service.GetUserByEmail(dto.Email)
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

// FetchUserHandler fetches user data by ID
func FetchUserHandler(w http.ResponseWriter, r *http.Request, service *services.UserService) {
	vars := mux.Vars(r)
	userIDString, exists := vars["id"]
	if !exists {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Convert userID to int
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetch user from the service layer
	user, err := service.GetUserByID(userID)
	if err != nil {
		log.Printf("Error fetching user %d: %v", userID, err)
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}


// Do a partial update to a user (just updating email or password for e.g)
func PatchUserHandler(w http.ResponseWriter, r *http.Request, service *services.UserService) {
	userID, err := utils.ParseID(r)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid or missing user ID",
		})
		return
	}
	var update dto.UpsertUser
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid request payload",
		})
		return
	}
	id := uint(userID)
	update.UserID = &id

	user, err := service.UpsertUser(update)
	if err != nil {
		if err.Error() == "user not found" {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "User not found",
			})
			return
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Internal server error	",
		})
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, utils.Response{
		Status:  "success",
		Data: user,
	})
}

// Delete a user
func DeleteUserHandler(w http.ResponseWriter, r *http.Request, service *services.UserService) {
	vars := mux.Vars(r)
	userIDString, exists := vars["id"]
	if !exists {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Convert userID to int
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetch user from the service layer
	user, err := service.GetUserByID(userID)
	if err != nil {
		log.Printf("Error fetching user %d: %v", userID, err)
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}