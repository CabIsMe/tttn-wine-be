package services

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
)

type ProductService interface {
	AllProductsService() ([]models.Product, *internal.SystemStatus)
	NewReleaseProductsService() ([]models.Product, *internal.SystemStatus)
	GetProductService(productId string) (*models.Product, *internal.SystemStatus)
	PromotionalProductsService() ([]models.Product, *internal.SystemStatus)
}
type product_service struct {
	rp repositories.Repos
}

func NewProductService(rp repositories.Repos) ProductService {
	return &product_service{
		rp,
	}
}

func (s *product_service) AllProductsService() ([]models.Product, *internal.SystemStatus) {
	listData, err := s.rp.GetAllProducts()
	if err != nil {
		return nil, internal.SysStatus.DbFailed
	}
	return listData, nil
}
func (s *product_service) NewReleaseProductsService() ([]models.Product, *internal.SystemStatus) {
	listData, err := s.rp.GetNewReleaseProducts()
	if err != nil {
		return nil, internal.SysStatus.DbFailed
	}
	return listData, nil
}
func (s *product_service) GetProductService(productId string) (*models.Product, *internal.SystemStatus) {
	result, err := s.rp.GetProduct(productId)
	if err != nil {
		return nil, internal.SysStatus.DbFailed
	}
	return result, nil
}
func (s *product_service) PromotionalProductsService() ([]models.Product, *internal.SystemStatus) {
	listData, err := s.rp.GetPromotionalProducts()
	if err != nil {
		return nil, internal.SysStatus.DbFailed
	}
	return listData, nil
}
