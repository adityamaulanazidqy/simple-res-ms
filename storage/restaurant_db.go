package storage

import (
	"restaurant/model"
)

type RestaurantDB struct {
	Order   []model.Order
	Product []model.Product
	User    []model.User
}

var restaurantDB = &RestaurantDB{}

func NewRestaurantDB() *RestaurantDB {
	return restaurantDB
}
