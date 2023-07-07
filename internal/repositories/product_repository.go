package repositories

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
}

type product_repos struct {
}

func NewProductRepository() ProductRepository {
	return &product_repos{}
}

func (r *product_repos) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := internal.Db.Debug().Find(&products)
	return products, err.Error
}
