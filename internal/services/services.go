package services

import "github.com/CabIsMe/tttn-wine-be/internal/repositories"

type Services struct {
	Rp repositories.Repos
	ProductService
}

func NewServices(rp repositories.Repos) Services {
	return Services{
		rp,
		NewProductService(rp),
	}
}
