package repositories

import (
	"fmt"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetTheTopSellingProducts() ([]models.Product, error)
	GetNewReleaseProducts() ([]models.Product, error)
}

type product_repos struct {
}

func NewProductRepository() ProductRepository {
	return &product_repos{}
}

func (r *product_repos) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	var model models.Product
	err := internal.Db.Debug().
		Select("product.*, brand.brand_name").
		Table(model.TableName()).
		Joins("JOIN brand ON product.brand_id = brand.brand_id").Preload("BrandInfo").
	Find(&products).Error
	return products, err
}

func (r *product_repos) GetTheTopSellingProducts() ([]models.Product, error) {
	var products []models.Product
	var model models.Product
	err := internal.Db.Debug().
		Select("product.*, brand.brand_name").
		Table(model.TableName()).
		Joins("JOIN brand ON product.brand_id = brand.brand_id").
		Find(&products)
	return products, err.Error
}

func (r *product_repos) GetNewReleaseProducts() ([]models.Product, error) {
	var products []models.Product
	var model models.Product
	err := internal.Db.Debug().Table(model.TableName()).
		Where(fmt.Sprintf("%s = ?", model.ColumnIsNew()), "1").
		Find(&products)
	return products, err.Error
}
