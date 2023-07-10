package delivery

import (
	"net/http"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/services"
	"github.com/CabIsMe/tttn-wine-be/internal/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ProductHandler interface {
	AllProductsHandler(ctx *fiber.Ctx) error
	NewReleaseProductsHandler(ctx *fiber.Ctx) error
	GetProductHandler(ctx *fiber.Ctx) error
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

func (h *product_handler) GetProductHandler(ctx *fiber.Ctx) error {
	body := &struct {
		ProductId string `json:"product_id" validate:"required"`
	}{}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusOK).JSON(models.Resp{
			Status: internal.CODE_WRONG_PARAMS,
			Msg:    internal.MSG_WRONG_PARAMS,
		})
	}
	internal.Log.Info("GetProductHandler", zap.Any("product_id", body))
	errs := utils.ValidateStruct(body)
	if errs != nil {
		return ctx.Status(http.StatusOK).JSON(models.Resp{
			Status: internal.CODE_WRONG_PARAMS,
			Msg:    internal.MSG_WRONG_PARAMS,
			Detail: utils.ShowErrors(errs),
		})
	}
	result, err := h.s.GetProductService(body.ProductId)
	if err != nil {
		return ctx.Status(200).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: result,
	})
}
