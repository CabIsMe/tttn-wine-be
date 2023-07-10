package delivery

import (
	"net/http"

	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	AllProductsHandler(ctx *fiber.Ctx) error
	NewReleaseProductsHandler(ctx *fiber.Ctx) error
}
type product_handler struct {
	s services.MainServices
}

func NewProductHandler(s services.MainServices) ProductHandler {
	return &product_handler{
		s,
	}
}
func (h *product_handler) AllProductsHandler(ctx *fiber.Ctx) error {
	results, err := h.s.AllProductsService()
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: results,
	})
}

func (h *product_handler) NewReleaseProductsHandler(ctx *fiber.Ctx) error {
	results, err := h.s.NewReleaseProductsService()
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: results,
	})
}
