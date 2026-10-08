package repository

import (
	"errors"

	"github.com/Raden-Salaf/E-commerce_Microservices/product-service/internal/model"
)

// ErrProductNotFound dikembalikan ketika produk tidak ada.
var ErrProductNotFound = errors.New("product not found")

// ProductRepository adalah kontrak akses data produk.
type ProductRepository interface {
	FindAll() []model.Product
	FindByID(id int) (model.Product, error)
}

type memoryProductRepository struct {
	products []model.Product
}

// NewMemoryProductRepository membuat repository dengan data di memori.
func NewMemoryProductRepository() ProductRepository {
	return &memoryProductRepository{
		products: []model.Product{
			{ID: 1, Name: "Keripik Singkong", Price: 15000, Stock: 40, CategoryID: 1},
			{ID: 2, Name: "Kopi Robusta 250g", Price: 45000, Stock: 25, CategoryID: 2},
			{ID: 3, Name: "Madu Hutan 500ml", Price: 85000, Stock: 10, CategoryID: 2},
		},
	}
}

func (r *memoryProductRepository) FindAll() []model.Product {
	return r.products
}

func (r *memoryProductRepository) FindByID(id int) (model.Product, error) {
	for _, p := range r.products {
		if p.ID == id {
			return p, nil
		}
	}
	return model.Product{}, ErrProductNotFound
}
