package delivery

import "github.com/CabIsMe/tttn-wine-be/internal/services"

type ProductHandler interface {
}
type product_handler struct {
	services.Services
}

func NewProductHandler(s services.Services) ProductHandler {
	return &product_handler{
		s,
	}
}
