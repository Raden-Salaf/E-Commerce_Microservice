package services

import (
	"errors"

	"github.com/Raden-Salaf/ecommerce-microservices/frontend/models"
)

var ErrProductNotFound = errors.New("product not found")

// Data tiruan. Nanti diganti dengan panggilan ke API Gateway.
var mockProducts = []models.Product{
	{ID: 1, Name: "Keripik Singkong", Price: 15000, Stock: 40, CategoryID: 1},
	{ID: 2, Name: "Kopi Robusta 250g", Price: 45000, Stock: 25, CategoryID: 2},
	{ID: 3, Name: "Madu Hutan 500ml", Price: 85000, Stock: 10, CategoryID: 2},
}

func GetProducts() ([]models.Product, error) {
	return mockProducts, nil
}

func GetProductByID(id int) (models.Product, error) {
	for _, p := range mockProducts {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Product{}, ErrProductNotFound
}
