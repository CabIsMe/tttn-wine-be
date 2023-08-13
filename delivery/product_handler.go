package delivery

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/services"
	"github.com/CabIsMe/tttn-wine-be/internal/utils"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
)

type ProductHandler interface {
	AllProductsHandler(ctx *fiber.Ctx) error
	NewReleaseProductsHandler(ctx *fiber.Ctx) error
	TopSellingProductsHandler(ctx *fiber.Ctx) error
	GetProductHandler(ctx *fiber.Ctx) error
	GetProductByNameHandler(ctx *fiber.Ctx) error
	PromotionalProductsHandler(ctx *fiber.Ctx) error
	GetProductsByTypeAndBrandHandler(ctx *fiber.Ctx) error
	GetProductsByBrandHandler(ctx *fiber.Ctx) error
	GetProductsByCategoryHandler(ctx *fiber.Ctx) error
	AddNewProductHandler(ctx *fiber.Ctx) error
}
type product_handler struct {
	services.MainServices
}

func NewProductHandler(s services.MainServices) ProductHandler {
	return &product_handler{
		s,
	}
}
func (h *product_handler) AllProductsHandler(ctx *fiber.Ctx) error {
	results, err := h.MainServices.ProductService.AllProductsService()
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
	results, err := h.MainServices.ProductService.TopSellingProductsService()
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
	results, err := h.MainServices.ProductService.NewReleaseProductsService()
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
	result, err := h.MainServices.ProductService.GetProductService(body.ProductId)
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
	result, err := h.MainServices.ProductService.GetProductsByNameService(body.ProductName)
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
	results, err := h.MainServices.ProductService.PromotionalProductsService()
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
		item.DiscountedCost = float32(math.Round(float64(item.Cost * (1 - item.PromotionDetailInfo.DiscountPercentage))))
		fmt.Println((1 - item.PromotionDetailInfo.DiscountPercentage), item.DiscountedCost)
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
func (h *product_handler) GetProductsByTypeAndBrandHandler(ctx *fiber.Ctx) error {
	resultError := models.Resp{
		Status: internal.CODE_WRONG_PARAMS,
		Msg:    internal.MSG_WRONG_PARAMS,
	}
	var body interface{}
	ctx.BodyParser(&body)
	uri := string(ctx.Request().URI().RequestURI())
	tokenAuth := string(ctx.Request().Header.Peek("token"))
	defer func() {
		internal.Log.Info("CreatePromotionHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth), zap.Any("body", body))
	}()
	payload := &struct {
		CustomerOrderId string `json:"customer_order_id" validate:"required"`
	}{}
	if err := ctx.BodyParser(&payload); err != nil {
		internal.Log.Error("BodyParser", zap.Any("Error", err.Error()))
		resultError.Detail = err.Error()
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	errs := utils.ValidateStruct(payload)
	if errs != nil {
		internal.Log.Error("ValidateStruct", zap.Any("Error", utils.ShowErrors(errs)))
		resultError.Detail = utils.ShowErrors(errs)
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	listData, errGet := h.MainServices.ProductService.GetProductsByTypeAndBrandService(payload.CustomerOrderId)
	if errGet != nil {
		ctx.Status(http.StatusOK).JSON(errGet)
	}
	output := make(map[string]interface{})
	output["listData"] = listData
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: output,
	})
}
func (h *product_handler) GetProductsByBrandHandler(ctx *fiber.Ctx) error {
	resultError := models.Resp{
		Status: internal.CODE_WRONG_PARAMS,
		Msg:    internal.MSG_WRONG_PARAMS,
	}
	var body interface{}
	ctx.BodyParser(&body)
	uri := string(ctx.Request().URI().RequestURI())
	tokenAuth := string(ctx.Request().Header.Peek("token"))
	defer func() {
		internal.Log.Info("GetProductsByBrandHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth), zap.Any("body", body))
	}()
	payload := &struct {
		BrandId string `json:"brand_id" validate:"required"`
	}{}
	if err := ctx.BodyParser(&payload); err != nil {
		internal.Log.Error("BodyParser", zap.Any("Error", err.Error()))
		resultError.Detail = err.Error()
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	errs := utils.ValidateStruct(payload)
	if errs != nil {
		internal.Log.Error("ValidateStruct", zap.Any("Error", utils.ShowErrors(errs)))
		resultError.Detail = utils.ShowErrors(errs)
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	listData, errGet := h.MainServices.ProductService.GetProductsByBrandService(payload.BrandId)
	if errGet != nil {
		ctx.Status(http.StatusOK).JSON(errGet)
	}
	output := make(map[string]interface{})
	output["listData"] = listData
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: output,
	})
}
func (h *product_handler) GetProductsByCategoryHandler(ctx *fiber.Ctx) error {
	resultError := models.Resp{
		Status: internal.CODE_WRONG_PARAMS,
		Msg:    internal.MSG_WRONG_PARAMS,
	}
	var body interface{}
	ctx.BodyParser(&body)
	uri := string(ctx.Request().URI().RequestURI())
	tokenAuth := string(ctx.Request().Header.Peek("token"))
	defer func() {
		internal.Log.Info("GetProductsByCategoryHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth), zap.Any("body", body))
	}()
	payload := &struct {
		CategoryId string `json:"category_id" validate:"required"`
	}{}
	if err := ctx.BodyParser(&payload); err != nil {
		internal.Log.Error("BodyParser", zap.Any("Error", err.Error()))
		resultError.Detail = err.Error()
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	errs := utils.ValidateStruct(payload)
	if errs != nil {
		internal.Log.Error("ValidateStruct", zap.Any("Error", utils.ShowErrors(errs)))
		resultError.Detail = utils.ShowErrors(errs)
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	listData, errGet := h.MainServices.ProductService.GetProductsByCategoryService(payload.CategoryId)
	if errGet != nil {
		ctx.Status(http.StatusOK).JSON(errGet)
	}
	output := make(map[string]interface{})
	output["listData"] = listData
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: output,
	})
}

func (h *product_handler) AddNewProductHandler(ctx *fiber.Ctx) error {
	resultError := models.Resp{
		Status: internal.CODE_WRONG_PARAMS,
		Msg:    internal.MSG_WRONG_PARAMS,
	}
	var body interface{}
	ctx.BodyParser(&body)
	uri := string(ctx.Request().URI().RequestURI())
	tokenAuth := string(ctx.Request().Header.Peek("token"))
	defer func() {
		internal.Log.Info("AddNewProductHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth), zap.Any("body", body))
	}()
	var payload *models.Product
	if err := ctx.BodyParser(&payload); err != nil {
		internal.Log.Error("BodyParser", zap.Any("Error", err.Error()))
		resultError.Detail = err.Error()
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	errs := utils.ValidateStruct(payload)
	if errs != nil {
		internal.Log.Error("ValidateStruct", zap.Any("Error", utils.ShowErrors(errs)))
		resultError.Detail = utils.ShowErrors(errs)
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	nanoId, _ := gonanoid.New()
	payload.ProductId = nanoId
	errCreate := h.MainServices.ProductService.AddNewProductService(*payload)
	if errCreate != nil {
		return ctx.Status(http.StatusOK).JSON(errCreate)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
	})
}
