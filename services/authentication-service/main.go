package main

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurant/model"
	"restaurant/storage"
)

func main() {
	restaurantDB := storage.NewRestaurantDB()

	mux := http.ServeMux{}
	mux.HandleFunc(
		"/login", func(w http.ResponseWriter, r *http.Request) {
			var request model.UserLoginRequest
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var foundUser *model.User
			for _, user := range restaurantDB.User {
				if user.Username == request.Username {
					foundUser = &user
					break
				}
			}

			if foundUser == nil {
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
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			existUser := false
			for _, user := range restaurantDB.User {
				if user.Username == request.Username {
					existUser = true
					break
				}
			}

			if existUser {
				http.Error(w, "User already exists", http.StatusConflict)
				return
			}

			restaurantDB.User = append(
				restaurantDB.User, model.User{
					ID:       len(restaurantDB.User) + 1,
					Username: request.Username,
					Password: request.Password,
				},
			)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				map[string]string{
					"message": "Registration successful",
				},
			)
		},
	)

	if err := http.ListenAndServe(":8081", &mux); err != nil {
		panic(err)
	}

	log.Println("Server started at :8080")
}
