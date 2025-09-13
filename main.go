package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"go-prisma-backend/routes"

	"github.com/joho/godotenv"
)

func main() {

	// // // Initialize Prisma client
	// client := db.NewClient()

	// // // Connect to Postgres
	// if err := client.Prisma.Connect(); err != nil {
	// 	log.Fatalf(" Could not connect to database: %v", err)
	// }
	// fmt.Println("✅ Connected to the database")
	// defer func() {
	// 	if err := client.Prisma.Disconnect(); err != nil {
	// 		log.Fatalf(" Failed to disconnect: %v", err)
	// 	}
	// }()




	// mapping the routtes of all controllers
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	// Option 1: Serve with prefix (/api/signup)
	http.Handle("/api/", http.StripPrefix("/api", mux))

	// Option 2: Serve without prefix (/signup)
	// http.Handle("/", mux)




	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println(" No .env file found, using system environment variables")
	}
	// Ensure DATABASE_URL is set
	if os.Getenv("DATABASE_URL") == "" {
		log.Fatal(" DATABASE_URL environment variable is not set")
	}



	// GET /users → fetch all users// // // Initialize Prisma client
	// client := db.NewClient()

	// // // Connect to Postgres
	// if err := client.Prisma.Connect(); err != nil {
	// 	log.Fatalf(" Could not connect to database: %v", err)
	// }
	// fmt.Println("✅ Connected to the database")
	// defer func() {
	// 	if err := client.Prisma.Disconnect(); err != nil {
	// 		log.Fatalf(" Failed to disconnect: %v", err)
	// 	}
	// }()
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json") // set JSON content type
		w.WriteHeader(http.StatusOK)                       // set HTTP status code

		response := map[string]string{
			"message": "Hello, World!",
			"status":  "success",
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	})





	// Catch-all handler for unmatched routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("❓ Unhandled request: %s %s", r.Method, r.URL.Path)
		http.NotFound(w, r)
	})




	fmt.Println("🚀 Server running at http://localhost:3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
