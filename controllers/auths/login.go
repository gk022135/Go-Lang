package auths

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5" // run: go get github.com/golang-jwt/jwt/v5
	"go-prisma-backend/prisma/db"
)

type LoginController struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type Response struct {
// 	Message string      `json:"message"`
// 	Status  bool        `json:"status"`
// 	Data    interface{} `json:"data,omitempty"`
// }

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login endpoint hit")

	// Only allow POST
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "method not allowed",
			Status:  false,
		})
		return
	}

	// Prisma client
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		http.Error(w, "failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer client.Prisma.Disconnect()

	// Decode request body
	var user LoginController
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "invalid request body",
			Status:  false,
		})
		return
	}

	// Find user by email
	existing, err := client.User.FindUnique(
		db.User.Email.Equals(user.Email),
	).Exec(context.Background())

	if err != nil || existing == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{
			Message: "invalid credentials",
			Status:  false,
		})
		return
	}

	// Check password (NOTE: you should use bcrypt.CompareHashAndPassword)
	if existing.Password != user.Password {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{
			Message: "invalid credentials",
			Status:  false,
		})
		return
	}

	// Generate JWT token
	secret := []byte(os.Getenv("JWT_SECRET")) // set in .env
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": existing.ID,
		"email":   existing.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // expires in 24h
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}

	// Set token as cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		HttpOnly: true,               // JS can’t read it
		Secure:   false,              // set true if using HTTPS
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
	})

	// Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Message: "login successful",
		Status:  true,
		Data: map[string]interface{}{
			"token": tokenString,
			"user":  existing,
		},
	})
}
