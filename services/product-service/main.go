package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restaurant/model"
	"restaurant/storage"
	"strconv"
)

func main() {
	restaurantDB := storage.NewRestaurantDB()

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/product", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				var product model.Product
				if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				product.ID = len(restaurantDB.Product) + 1
				restaurantDB.Product = append(restaurantDB.Product, product)

				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(
					map[string]interface{}{
						"message": "product created successfully",
						"product": product,
					},
				)
			} else if r.Method == http.MethodGet {
				if len(restaurantDB.Product) == 0 {
					w.WriteHeader(http.StatusNoContent)
					json.NewEncoder(w).Encode(map[string]string{"message": "no products found"})
					return
				}

				for _, product := range restaurantDB.Product {
					fmt.Println("ID\tNAME\tDESCRIPTION\tPRICE")
					fmt.Printf("%d\t%s\t%s\t%.2f\n", product.ID, product.Name, product.Description, product.Price)
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(restaurantDB.Product)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
		},
	)

	mux.HandleFunc(
		"/product/", func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path[len("/product/"):]
			id, err := strconv.Atoi(path)
			if err != nil {
				http.Error(w, "Invalid product ID", http.StatusBadRequest)
				return
			}

			switch r.Method {
			case http.MethodGet:
				var foundProduct *model.Product
				for i := range restaurantDB.Product {
					if restaurantDB.Product[i].ID == id {
						foundProduct = &restaurantDB.Product[i]
						break
					}
				}

				if foundProduct == nil {
					http.Error(w, "Product not found", http.StatusNotFound)
					return
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(foundProduct)

			case http.MethodPut:
				var updatedProduct model.Product
				if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				productIndex := -1
				for i := range restaurantDB.Product {
					if restaurantDB.Product[i].ID == id {
						productIndex = i
						break
					}
				}

				if productIndex == -1 {
					http.Error(w, "Product not found", http.StatusNotFound)
					return
				}

				updatedProduct.ID = id
				restaurantDB.Product[productIndex] = updatedProduct

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(
					map[string]interface{}{
						"message": "product updated successfully",
						"product": updatedProduct,
					},
				)

			case http.MethodDelete:
				productIndex := -1
				for i := range restaurantDB.Product {
					if restaurantDB.Product[i].ID == id {
						productIndex = i
						break
					}
				}

				if productIndex == -1 {
					http.Error(w, "Product not found", http.StatusNotFound)
					return
				}

				restaurantDB.Product = append(
					restaurantDB.Product[:productIndex],
					restaurantDB.Product[productIndex+1:]...,
				)

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{"message": "product deleted successfully"})

			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		},
	)

	if err := http.ListenAndServe(":8082", mux); err != nil {
		panic(err)
	}
}
