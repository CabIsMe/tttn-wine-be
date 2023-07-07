package delivery

import "github.com/CabIsMe/tttn-wine-be/internal/services"

type Handlers interface {
	ProductHandler
}
type handlers struct {
	ProductHandler
}

func NewHandlers(services services.Services) Handlers {
	return &handlers{
		services,
		NewProductHandler(services),
	}
}
