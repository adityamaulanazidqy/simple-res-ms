package storage

import (
	"log"
	"restaurant/model"
)

type OrderStorage struct {
	Orders []model.Order
}

var orderStorage *OrderStorage

func init() {
	orderStorage = &OrderStorage{
		Orders: make([]model.Order, 0),
	}
	log.Println("Order storage initialized with empty order list")
}

func NewOrderStorage() *OrderStorage {
	return orderStorage
}

func (s *OrderStorage) GetOrderByID(id int) (*model.Order, bool) {
	for i := range s.Orders {
		if s.Orders[i].ID == id {
			return &s.Orders[i], true
		}
	}
	return nil, false
}

func (s *OrderStorage) AddOrder(order model.Order) {
	s.Orders = append(s.Orders, order)
	log.Printf("Order added: ID=%d, UserID=%d, ProductID=%d", order.ID, order.UserID, order.ProductID)
}

func (s *OrderStorage) GetOrdersByUserID(userID int) []model.Order {
	var userOrders []model.Order
	for _, order := range s.Orders {
		if order.UserID == userID {
			userOrders = append(userOrders, order)
		}
	}
	return userOrders
}

func (s *OrderStorage) GetOrderCount() int {
	return len(s.Orders)
}
