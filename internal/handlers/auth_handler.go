package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"tasko/internal/models"
	"tasko/internal/repo"
	"tasko/pkg/auth"
)

func LoginHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginRequest models.UserCredentials

		if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		// validate user
		user, err := repo.GetUserByEmail(db, loginRequest.Email)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "User not found", http.StatusUnauthorized)
			} else {
				log.Println("Error retrieving user:", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
		if err != nil {
			http.Error(w, "Invalid password", http.StatusForbidden)
			return
		}

		// Generate JWT token
		token, err := auth.GenerateToken(int(user.ID))
		if err != nil {
			log.Println("Error generating token:", err)
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		// Return response
		response := struct {
			Token string `json:"token"`
		}{
			Token: token,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
