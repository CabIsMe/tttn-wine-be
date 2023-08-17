package services

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
)

type ProductService interface {
	AllProductsService() ([]models.Product, *internal.SystemStatus)
	TopSellingProductsService() ([]models.ProductAndFrequency, *internal.SystemStatus)
	NewReleaseProductsService() ([]models.Product, *internal.SystemStatus)
	GetProductService(productId string) (*models.Product, *internal.SystemStatus)
	GetProductsByNameService(productId string) ([]models.Product, *internal.SystemStatus)
	PromotionalProductsService() ([]models.Product, *internal.SystemStatus)
	GetProductsByTypeAndBrandService(productId string) ([]models.Product, *internal.SystemStatus)
	GetProductsByBrandService(productId string) ([]models.Product, *internal.SystemStatus)
	GetProductsByCategoryService(productId string) ([]models.Product, *internal.SystemStatus)
	AddNewProductService(product models.Product) *internal.SystemStatus
	UpdateProductService(product models.Product) *internal.SystemStatus
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
func (s *product_service) TopSellingProductsService() ([]models.ProductAndFrequency, *internal.SystemStatus) {
	listData, err := s.rp.GetTopSellingProducts()
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
func (s *product_service) GetProductsByNameService(productName string) ([]models.Product, *internal.SystemStatus) {
	results, err := s.rp.GetProductsByName(productName)
	if err != nil {
		return nil, internal.SysStatus.DbFailed
	}
	return results, nil
}
func (s *product_service) PromotionalProductsService() ([]models.Product, *internal.SystemStatus) {
	listData, err := s.rp.GetPromotionalProducts()
	if err != nil {
		return nil, internal.SysStatus.DbFailed
	}
	return listData, nil
}
func (s *product_service) GetProductsByTypeAndBrandService(productId string) ([]models.Product, *internal.SystemStatus) {
	listData, err := s.rp.GetProductsByTypeAndBrand(productId)
	if err != nil {
		return nil, internal.SysStatus.DbFailed
	}
	return listData, nil
}
func (s *product_service) GetProductsByBrandService(brandId string) ([]models.Product, *internal.SystemStatus) {
	listData, err := s.rp.GetProductsByBrand(brandId)
	if err != nil {
		return nil, internal.SysStatus.DbFailed
	}
	return listData, nil
}
func (s *product_service) GetProductsByCategoryService(categoryId string) ([]models.Product, *internal.SystemStatus) {
	listData, err := s.rp.GetProductsByCategory(categoryId)
	if err != nil {
		return nil, internal.SysStatus.DbFailed
	}
	return listData, nil
}
func (s *product_service) AddNewProductService(product models.Product) *internal.SystemStatus {

	err := s.rp.AddNewProduct(product)
	if err != nil {
		return internal.SysStatus.DbFailed
	}
	return nil
}
func (s *product_service) UpdateProductService(product models.Product) *internal.SystemStatus {

	err := s.rp.UpdateProduct(product)
	if err != nil {
		return internal.SysStatus.DbFailed
	}
	return nil
}
