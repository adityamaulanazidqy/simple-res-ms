package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restaurant/model"
	"restaurant/storage"
)

func main() {
	restaurantDB := storage.NewRestaurantDB()

	mux := http.NewServeMux()
	mux.HandleFunc(
		"/order", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				var orderRequest model.OrderRequest
				if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				existUser := false
				existProduct := false
				for _, user := range restaurantDB.User {
					if user.ID == orderRequest.UserID {
						existUser = true
						break
					}
				}

				for _, product := range restaurantDB.Product {
					if product.ID == orderRequest.ProductID {
						existProduct = true
						break
					}
				}

				if !existUser {
					http.Error(w, "User not found", http.StatusNotFound)
					return
				}

				if !existProduct {
					http.Error(w, "Product not found", http.StatusNotFound)
					return
				}

				restaurantDB.Order = append(
					restaurantDB.Order, model.Order{
						ID:         len(restaurantDB.Order) + 1,
						UserID:     orderRequest.UserID,
						ProductID:  orderRequest.ProductID,
						Quantity:   orderRequest.Quantity,
						TotalPrice: orderRequest.TotalPrice,
					},
				)

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(
					map[string]interface{}{
						"message": "order placed successfully",
						"order":   orderRequest,
					},
				)
			} else if r.Method == http.MethodGet {
				if len(restaurantDB.Order) == 0 {
					w.WriteHeader(http.StatusNoContent)
					json.NewEncoder(w).Encode(map[string]string{"message": "no orders found"})
					return
				}

				for _, order := range restaurantDB.Order {
					fmt.Println("ID\tUSERID\tPRODUCTID\tQUANTITY\tTOTAL PRICE")
					fmt.Printf(
						"%d\t%d\t%d\t%d\t%.2f\n",
						order.ID,
						order.UserID,
						order.ProductID,
						order.Quantity,
						order.TotalPrice,
					)
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(restaurantDB.Order)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
		},
	)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
