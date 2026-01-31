package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restaurant/model"
	"restaurant/storage"
)

func main() {
	orderDB := storage.NewOrderStorage()

	if orderDB.GetOrderCount() == 0 {
		orderDB.AddOrder(model.Order{ID: 1, UserID: 1, ProductID: 1, Quantity: 2, TotalPrice: 31.98})
		orderDB.AddOrder(model.Order{ID: 2, UserID: 2, ProductID: 3, Quantity: 1, TotalPrice: 8.99})
	}

	var checkUserFound = func(userID int) (
		error,
		int,
	) {
		request, err := http.NewRequest(
			http.MethodGet,
			fmt.Sprintf("http://auth-service:8081/users/%d", userID),
			nil,
		)

		if err != nil {
			return fmt.Errorf("error creating request"), http.StatusInternalServerError
		}

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return fmt.Errorf("error checking user existence"), http.StatusInternalServerError
		}
		defer response.Body.Close()

		if response.StatusCode == http.StatusNotFound {
			return fmt.Errorf("user not found"), http.StatusNotFound
		}

		if response.StatusCode != http.StatusOK {
			return fmt.Errorf("error checking user existence"), http.StatusInternalServerError
		}

		return nil, http.StatusOK
	}

	var checkProductFound = func(productID int) (
		error,
		int,
	) {
		url := fmt.Sprintf("http://product-service:8082/product/%d", productID)
		log.Printf("Checking product existence: %s", url)
		request, err := http.NewRequest(
			http.MethodGet,
			url,
			nil,
		)

		if err != nil {
			log.Printf("Error creating product request: %v", err)
			return fmt.Errorf("error creating request"), http.StatusInternalServerError
		}

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			log.Printf("Error making product request: %v", err)
			return fmt.Errorf("product not found"), http.StatusNotFound
		}
		defer response.Body.Close()

		log.Printf("Product response status: %d", response.StatusCode)
		if response.StatusCode == http.StatusNotFound {
			return fmt.Errorf("product not found"), http.StatusNotFound
		}

		if response.StatusCode != http.StatusOK {
			return fmt.Errorf("error checking product existence"), http.StatusInternalServerError
		}

		return nil, http.StatusOK
	}

	mux := http.NewServeMux()
	mux.HandleFunc(
		"/order", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				var orderRequest model.OrderRequest
				if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
					log.Printf("Error decoding order request: %v", err)
					http.Error(w, "Invalid request body", http.StatusBadRequest)
					return
				}

				log.Printf("Creating order for user %d, product %d", orderRequest.UserID, orderRequest.ProductID)

				if err, status := checkUserFound(orderRequest.UserID); err != nil {
					http.Error(w, err.Error(), status)
					return
				}

				if err, status := checkProductFound(orderRequest.ProductID); err != nil {
					http.Error(w, err.Error(), status)
					return
				}

				orderDB.AddOrder(model.Order{
					ID:         orderDB.GetOrderCount() + 1,
					UserID:     orderRequest.UserID,
					ProductID:  orderRequest.ProductID,
					Quantity:   orderRequest.Quantity,
					TotalPrice: orderRequest.TotalPrice,
				})

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(
					map[string]interface{}{
						"message": "order placed successfully",
						"order":   orderRequest,
					},
				)
			} else if r.Method == http.MethodGet {
				if len(orderDB.Orders) == 0 {
					w.WriteHeader(http.StatusNoContent)
					json.NewEncoder(w).Encode(map[string]string{"message": "no orders found"})
					return
				}

				for _, order := range orderDB.Orders {
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
				json.NewEncoder(w).Encode(orderDB.Orders)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
		},
	)

	log.Println("Order service starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
