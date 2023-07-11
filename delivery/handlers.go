package delivery

import (
	"github.com/CabIsMe/tttn-wine-be/internal/services"
	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	// checkRequiredFields(requiredFields []string) fiber.Handler
	VerifyToken(ctx *fiber.Ctx) error
	ProductHandler
	AuthenticationHandler
	PromotionHandler
}
type handlers struct {
	services.MainServices
	ProductHandler
	AuthenticationHandler
	PromotionHandler
}

func NewHandlers(s services.MainServices) Handlers {
	return &handlers{
		s,
		NewProductHandler(s),
		NewAuthenticationHandler(s),
		NewPromotionHandler(s),
	}
}
