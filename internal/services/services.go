package services

import "github.com/CabIsMe/tttn-wine-be/internal/repositories"

type MainServices struct {
	repositories.Repos
	ProductService
	AuthenticationService
	PromotionService
	CustomerOrderService
	CustomerService
}

func NewServices(rp repositories.Repos) MainServices {
	return MainServices{
		rp,
		NewProductService(rp),
		NewAuthenticationService(rp),
		NewPromotionService(rp),
		NewCustomerOrderService(rp),
		NewCustomerService(rp),
	}
}
