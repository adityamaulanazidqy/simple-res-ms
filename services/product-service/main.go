package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restaurant/model"
	"restaurant/storage"
	"strconv"
)

func main() {
	productDB := storage.NewProductStorage()

	if productDB.GetProductCount() == 0 {
		productDB.AddProduct(model.Product{ID: 1, Name: "Burger", Description: "Delicious beef burger", Price: 15.99})
		productDB.AddProduct(model.Product{ID: 2, Name: "Pizza", Description: "Margherita pizza", Price: 12.50})
		productDB.AddProduct(model.Product{ID: 3, Name: "Salad", Description: "Fresh garden salad", Price: 8.99})
	}

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/product", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				var product model.Product
				if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
					log.Printf("Error decoding product: %v", err)
					http.Error(w, "Invalid request body", http.StatusBadRequest)
					return
				}

				log.Printf("Creating product: %s", product.Name)

				product.ID = productDB.GetProductCount() + 1
				productDB.AddProduct(product)

				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(
					map[string]interface{}{
						"message": "product created successfully",
						"product": product,
					},
				)
			} else if r.Method == http.MethodGet {
				if len(productDB.Products) == 0 {
					w.WriteHeader(http.StatusNoContent)
					json.NewEncoder(w).Encode(map[string]string{"message": "no products found"})
					return
				}

				for _, product := range productDB.Products {
					fmt.Println("ID\tNAME\tDESCRIPTION\tPRICE")
					fmt.Printf("%d\t%s\t%s\t%.2f\n", product.ID, product.Name, product.Description, product.Price)
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(productDB.Products)
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
				foundProduct, exists := productDB.GetProductByID(id)
				if !exists {
					http.Error(w, "Product not found", http.StatusNotFound)
					return
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(foundProduct)

			case http.MethodPut:
				var updatedProduct model.Product
				if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
					log.Printf("Error decoding update product: %v", err)
					http.Error(w, "Invalid request body", http.StatusBadRequest)
					return
				}

				log.Printf("Updating product ID: %d", id)

				if !productDB.UpdateProduct(id, updatedProduct) {
					http.Error(w, "Product not found", http.StatusNotFound)
					return
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(
					map[string]interface{}{
						"message": "product updated successfully",
						"product": updatedProduct,
					},
				)

			case http.MethodDelete:
				if !productDB.DeleteProduct(id) {
					http.Error(w, "Product not found", http.StatusNotFound)
					return
				}

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
