package delivery

import (
	"net/http"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/services"
	"github.com/CabIsMe/tttn-wine-be/internal/utils"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
)

type PromotionHandler interface {
	CreatePromotionHandler(ctx *fiber.Ctx) error
	CreatePromotionDetailHandler(ctx *fiber.Ctx) error
	//TODO: get promotion by date and max(discount)
	GetPromotionByDateHandler(ctx *fiber.Ctx) error
}
type promotion_handler struct {
	services.MainServices
}

func NewPromotionHandler(s services.MainServices) PromotionHandler {
	return &promotion_handler{
		s,
	}
}
func (h *promotion_handler) CreatePromotionHandler(ctx *fiber.Ctx) error {
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
	var payload models.PromotionInput
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
	nanoId, err := gonanoid.New()
	if err != nil {
		internal.Log.Error("gonanoid", zap.Any("Error", err.Error()))
		resultError.Detail = err
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	dateStart, errParse := utils.ParseTimeFrStringV2("Y-M-D", payload.DateStart)
	if errParse != nil {
		internal.Log.Error("dateStart", zap.Any("Error", errParse.Error()))
		resultError.Detail = errParse.Error()
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	dateEnd, errParse := utils.ParseTimeFrStringV2("Y-M-D", payload.DateEnd)
	if errParse != nil {
		internal.Log.Error("dateEnd", zap.Any("Error", errParse.Error()))
		resultError.Detail = errParse.Error()
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	employeeId, ok := ctx.Locals("user_id").(string)
	if !ok {
		internal.Log.Error("employeeId, ok", zap.Any("Error", ok))
		return ctx.Status(http.StatusOK).JSON(resultError)
	}

	promotion := &models.Promotion{
		PromotionId:   nanoId,
		PromotionName: payload.PromotionName,
		DateStart:     dateStart,
		DateEnd:       dateEnd,
		Description:   payload.Description,
		EmployeeId:    employeeId,
	}

	errCreate := h.MainServices.PromotionService.CreatePromotionService(*promotion)
	if errCreate != nil {
		return ctx.Status(http.StatusOK).JSON(errCreate)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: payload,
	})

}
func (h *promotion_handler) CreatePromotionDetailHandler(ctx *fiber.Ctx) error {
	resultError := models.Resp{
		Status: internal.CODE_WRONG_PARAMS,
		Msg:    internal.MSG_WRONG_PARAMS,
	}
	var body interface{}
	ctx.BodyParser(&body)
	uri := string(ctx.Request().URI().RequestURI())
	tokenAuth := string(ctx.Request().Header.Peek("token"))
	defer func() {
		internal.Log.Info("CreatePromotionDetailHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth), zap.Any("body", body))
	}()
	var payload models.PromotionDetail
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
	errCreate := h.MainServices.PromotionService.CreatePromotionDetailService(payload)
	if errCreate != nil {
		internal.Log.Error("CreatePromotionDetailService", zap.Any("Error", errCreate))
		return ctx.Status(http.StatusOK).JSON(errCreate)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: payload,
	})
}

func (h *promotion_handler) GetPromotionByDateHandler(ctx *fiber.Ctx) error {
	uri := string(ctx.Request().URI().RequestURI())
	tokenAuth := string(ctx.Request().Header.Peek("token"))
	defer func() {
		internal.Log.Info("CreatePromotionDetailHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth))
	}()
	model, err := h.MainServices.PromotionService.GetPromotionByDateService()
	if err != nil {
		return ctx.Status(http.StatusOK).JSON(err)
	}
	if model == nil {
		return ctx.Status(http.StatusOK).JSON(models.Resp{
			Status: 1,
			Msg:    "OK",
			Detail: nil,
		})
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: model,
	})
}
