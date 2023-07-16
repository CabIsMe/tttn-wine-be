package delivery

import (
	"github.com/CabIsMe/tttn-wine-be/internal/services"
	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	// checkRequiredFields(requiredFields []string) fiber.Handler
	VerifyTokenClient(ctx *fiber.Ctx) error
	VerifyTokenInside(ctx *fiber.Ctx) error
	ProductHandler
	AuthenticationHandler
	PromotionHandler
	CustomerOrderHandler
}
type handlers struct {
	services.MainServices
	ProductHandler
	AuthenticationHandler
	PromotionHandler
	CustomerOrderHandler
}

func NewHandlers(s services.MainServices) Handlers {
	return &handlers{
		s,
		NewProductHandler(s),
		NewAuthenticationHandler(s),
		NewPromotionHandler(s),
		NewCustomerOrderHandler(s),
	}
}
