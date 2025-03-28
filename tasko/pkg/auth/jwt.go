// Package auth
// Author: Samoei Oloo
// Created: 2025-03-28
// License: None
//
// This script is responsible for JWT generation
package auth

import (
	"errors"
	"log"
	"os"
	"time"

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

// ParseJWT parses the JWT token
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

// checks if the JWT token is valid
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
