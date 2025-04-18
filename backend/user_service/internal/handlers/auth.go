package handlers

import (
	"encoding/json"
	"net/http"

	utils "github.com/samoei-ftw/specno/backend/common/utils"
	"github.com/samoei-ftw/specno/backend/user_service/internal/services"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, service *services.UserService) {
	var registerDTO struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&registerDTO); err != nil {
		utils.RespondWithErrorMessage(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if registerDTO.Email == "" || registerDTO.Password == "" {
		utils.RespondWithErrorMessage(w, http.StatusBadRequest, "Email and password are required")
		return
	}

	userID, err := service.RegisterUser(registerDTO.Email, registerDTO.Password)
	if err != nil {
		utils.RespondWithErrorMessage(w, http.StatusInternalServerError, "Registration failed")
		return
	}

	utils.RespondWithSuccess(w, http.StatusCreated, map[string]interface{}{
		"id":      userID,
		"message": "User registered successfully",
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request, service *services.UserService) {
	var loginDTO struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginDTO); err != nil {
		utils.RespondWithErrorMessage(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := service.AuthenticateUser(loginDTO.Email, loginDTO.Password)
	if err != nil {
		utils.RespondWithErrorMessage(w, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := utils.GenerateToken(int(user.ID))
	if err != nil {
		utils.RespondWithErrorMessage(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.RespondWithSuccess(w, http.StatusOK, map[string]interface{}{
		"token":   token,
		"role":    user.Role,
		"user_id": int(user.ID),
	})
}