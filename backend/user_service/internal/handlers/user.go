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
)

func FetchUserHandler(w http.ResponseWriter, r *http.Request, service *services.UserService) {
	vars := mux.Vars(r)
	userIDString, exists := vars["id"]
	if !exists {
		utils.RespondWithErrorMessage(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// Convert userID to int
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		utils.RespondWithErrorMessage(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Fetch user from the service layer
	user, err := service.GetUserByID(userID)
	if err != nil {
		log.Printf("Error fetching user %d: %v", userID, err)
		if err.Error() == "user not found" {
			utils.RespondWithErrorMessage(w, http.StatusNotFound, "User not found")
			return
		}
		utils.RespondWithErrorMessage(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	utils.RespondWithSuccess(w, http.StatusOK, user)
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
	update.UserID = utils.UintPtr(uint(userID))
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
			Status:  "error",
			Message: "Invalid request payload",
		})
		return
	}

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

func DeleteUserHandler(w http.ResponseWriter, r *http.Request, service *services.UserService) {
	vars := mux.Vars(r)
	userIDString, exists := vars["id"]
	if !exists {
		utils.RespondWithErrorMessage(w, http.StatusBadRequest, "User ID is required")
		return
	}

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		utils.RespondWithErrorMessage(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	claims, ok := r.Context().Value(utils.ClaimsKey).(*utils.Claims)
	if !ok {
		utils.RespondWithErrorMessage(w, http.StatusUnauthorized, "User not authorized")
		return
	}
	isDeleted, err := service.DeleteUser(userID, claims.UserID)
	if err != nil {
		log.Printf("Error deleting user %d: %v", userID, err)
		if err.Error() == "user not found" {
			utils.RespondWithErrorMessage(w, http.StatusNotFound, "User not found")
			return
		} else if err.Error() == "unauthorized: only admins can delete users" {
			utils.RespondWithErrorMessage(w, http.StatusForbidden, "Forbidden: Only admins can delete users")
			return
		}
		utils.RespondWithErrorMessage(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	utils.RespondWithSuccess(w, http.StatusOK, map[string]bool{"deleted": isDeleted})
}