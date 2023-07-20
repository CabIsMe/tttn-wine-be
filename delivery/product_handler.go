package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

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
	TopSellingProductsHandler(ctx *fiber.Ctx) error
	GetProductHandler(ctx *fiber.Ctx) error
	GetProductByNameHandler(ctx *fiber.Ctx) error
	PromotionalProductsHandler(ctx *fiber.Ctx) error
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
func (h *product_handler) TopSellingProductsHandler(ctx *fiber.Ctx) error {
	results, err := h.s.TopSellingProductsService()
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
func (h *product_handler) GetProductByNameHandler(ctx *fiber.Ctx) error {
	body := &struct {
		ProductName string `json:"product_name" validate:"required"`
	}{}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusOK).JSON(models.Resp{
			Status: internal.CODE_WRONG_PARAMS,
			Msg:    internal.MSG_WRONG_PARAMS,
		})
	}
	internal.Log.Info("GetProductByNameHandler", zap.Any("product_name", body))
	errs := utils.ValidateStruct(body)
	if errs != nil {
		return ctx.Status(http.StatusOK).JSON(models.Resp{
			Status: internal.CODE_WRONG_PARAMS,
			Msg:    internal.MSG_WRONG_PARAMS,
			Detail: utils.ShowErrors(errs),
		})
	}
	result, err := h.s.GetProductsByNameService(body.ProductName)
	if err != nil {
		return ctx.Status(200).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: result,
	})
}
func (h *product_handler) PromotionalProductsHandler(ctx *fiber.Ctx) error {
	results, err := h.s.PromotionalProductsService()
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(err)
	}
	jsonData, errMarshal := json.Marshal(results)
	if errMarshal != nil {
		fmt.Println("Error:", errMarshal.Error())
		return ctx.Status(fiber.StatusOK).JSON(errMarshal)
	}
	type promotionOutput struct {
		ProductId           string                 `json:"product_id"`
		ProductName         string                 `json:"product_name"`
		Cost                float32                `json:"cost"`
		ProductImg          string                 `json:"product_img"`
		Description         string                 `json:"description"`
		InventoryNumber     int                    `json:"inventory_number"`
		Status              string                 `json:"status"`
		BrandId             string                 `json:"-"`
		CategoryId          string                 `json:"-"`
		IsNew               int8                   `json:"is_new"`
		BrandInfo           models.Brand           `json:"brand_info"`
		CategoryInfo        models.Category        `json:"category_info"`
		PromotionDetailInfo models.PromotionDetail `json:"promotion_detail_info"`
		DiscountedCost      float32                `json:"discounted_cost"`
	}
	var data []*promotionOutput
	errMarshal = json.Unmarshal(jsonData, &data)
	if errMarshal != nil {
		fmt.Println("Error:", errMarshal.Error())
		return ctx.Status(fiber.StatusOK).JSON(errMarshal)
	}
	for _, item := range data {
		item.DiscountedCost = item.Cost * (1 - item.PromotionDetailInfo.DiscountPercentage)
		fmt.Println(item.DiscountedCost)
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].PromotionDetailInfo.DiscountPercentage*100 > data[j].PromotionDetailInfo.DiscountPercentage*100
	})
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: data,
	})
}
