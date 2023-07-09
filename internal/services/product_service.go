package services

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
)

type ProductService interface {
	AllProductsService() ([]models.Product, *internal.SystemStatus)
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
