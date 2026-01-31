package storage

import (
	"log"
	"restaurant/model"
)

type ProductStorage struct {
	Products []model.Product
}

var productStorage *ProductStorage

func init() {
	productStorage = &ProductStorage{
		Products: make([]model.Product, 0),
	}
	log.Println("Product storage initialized with empty product list")
}

func NewProductStorage() *ProductStorage {
	return productStorage
}

func (s *ProductStorage) GetProductByID(id int) (*model.Product, bool) {
	for i := range s.Products {
		if s.Products[i].ID == id {
			return &s.Products[i], true
		}
	}
	return nil, false
}

func (s *ProductStorage) AddProduct(product model.Product) {
	s.Products = append(s.Products, product)
	log.Printf("Product added: ID=%d, Name=%s", product.ID, product.Name)
}

func (s *ProductStorage) UpdateProduct(id int, product model.Product) bool {
	for i := range s.Products {
		if s.Products[i].ID == id {
			product.ID = id 
			s.Products[i] = product
			log.Printf("Product updated: ID=%d, Name=%s", id, product.Name)
			return true
		}
	}
	return false
}

func (s *ProductStorage) DeleteProduct(id int) bool {
	for i := range s.Products {
		if s.Products[i].ID == id {
			s.Products = append(s.Products[:i], s.Products[i+1:]...)
			log.Printf("Product deleted: ID=%d", id)
			return true
		}
	}
	return false
}

func (s *ProductStorage) GetProductCount() int {
	return len(s.Products)
}
