package services

import "github.com/CabIsMe/tttn-wine-be/internal/repositories"

type ProductService interface {
}
type product_service struct {
}

func NewProductService(rp repositories.Repos) ProductService {
	return &product_service{}
}
