package main

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurant/model"
	"restaurant/storage"
	"strconv"
)

func main() {
	userDB := storage.NewUserStorage()

	if userDB.GetUserCount() == 0 {
		userDB.AddUser(model.User{ID: 1, Username: "admin", Password: "admin123"})
		userDB.AddUser(model.User{ID: 2, Username: "user1", Password: "password123"})
	}

	mux := http.ServeMux{}
	mux.HandleFunc(
		"/login", func(w http.ResponseWriter, r *http.Request) {
			var request model.UserLoginRequest
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				log.Printf("Error decoding login request: %v", err)
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			log.Printf("Login attempt for user: %s", request.Username)

			foundUser, exists := userDB.GetUserByUsername(request.Username)
			if !exists {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}

			if foundUser.Password != request.Password {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]string{
					"message": "Login successful",
					"token":   "some-jwt-token",
				},
			)
		},
	)

	mux.HandleFunc(
		"/register", func(w http.ResponseWriter, r *http.Request) {
			var request model.UserRegisterRequest
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				log.Printf("Error decoding register request: %v", err)
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			log.Printf("Register attempt for user: %s", request.Username)

			if userDB.UserExists(request.Username) {
				http.Error(w, "User already exists", http.StatusConflict)
				return
			}

			userDB.AddUser(model.User{
				ID:       userDB.GetUserCount() + 1,
				Username: request.Username,
				Password: request.Password,
			})

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]string{
					"message": "Registration successful",
				},
			)
		},
	)

	mux.HandleFunc(
		"/users", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			if userDB.GetUserCount() == 0 {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNoContent)
				json.NewEncoder(w).Encode(map[string]string{"message": "no users found"})
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(userDB.Users)
		},
	)

	mux.HandleFunc(
		"/users/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			strID := r.URL.Path[len("/users/"):]
			id, err := strconv.Atoi(strID)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusBadRequest)
				return
			}

			foundUser, exists := userDB.GetUserByID(id)
			if !exists {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "user found", "username": foundUser.Username})
		},
	)

	log.Println("Authentication service starting on :8081")
	if err := http.ListenAndServe(":8081", &mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
