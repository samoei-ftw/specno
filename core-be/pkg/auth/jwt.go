// Package auth
// Author: Samoei Oloo
// Created: 2025-03-28
// License: None
//
// This script is responsible for JWT generation
package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"tasko/internal/models"
	"time"

	//"github.com/dgrijalva/jwt-go"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var jwtKey []byte

func init() {
	//load env
	err := godotenv.Load()

	if err != nil {
		log.Println("Warning")
	}

	// get jwt secret key
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET not set in environment variables")
	}
	jwtKey = []byte(secret)
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token
func GenerateToken(userID int) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

func ValidateToken(tokenStr string) (bool, error) {
	claims, err := ParseJWT(tokenStr)
	if err != nil {
		return false, err
	}

	// Check expiration
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return false, errors.New("token has expired")
	}

	return true, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.UserCredentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate user credentials
	/**user, err := repo.GetUserByEmail(credentials.Email)
	  if err != nil || user.Password != credentials.Password {
	      http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	      return
	  }*/

	// Generate JWT token
	token, err := GenerateToken(user)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
		tokenString = strings.Split(tokenString, " ")[1]
		/** token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		        return nil, fmt.Errorf("invalid signing method")
		    }
		    return jwtSecretKey, nil
		})*/

		/**  if err != nil || !token.Valid {
		    http.Error(w, "Invalid token", http.StatusUnauthorized)
		    return
		}*/

		// Proceed to the next handler if the token is valid
		next.ServeHTTP(w, r)
	})
}
