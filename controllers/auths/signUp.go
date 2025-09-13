package auths

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go-prisma-backend/prisma/db"
)

type SignUpController struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Response struct {
	Message  string      `json:"message"`
	Status   bool        `json:"status"`
	TryLater string      `json:"try_later,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}



func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signup endpoint hit")

	// Only allow POST method
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "method not allowed",
			Status:  false,
		})
		return
	}

	// Create Prisma client and connect
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		http.Error(w, "failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer client.Prisma.Disconnect()

	// Decode user JSON
	var user SignUpController
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Message: "invalid request body",
			Status:  false,
		})
		return
	}

	// Save to database
	newUser, err := client.User.CreateOne(
		db.User.FirstName.Set(user.FirstName),
		db.User.LastName.Set(user.LastName),
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.Password),
	).Exec(context.Background())

	if err != nil {
		fmt.Println("Error creating user:", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Message:  "failed to create user",
			Status:   false,
			TryLater: "please try again later",
		})
		return
	}

	fmt.Println("new user response:", newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Message: "user created successfully",
		Status:  true,
		Data:    newUser,
	})
}
