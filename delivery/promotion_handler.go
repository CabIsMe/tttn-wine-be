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
		resultError.Detail = err.Error()
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	errs := utils.ValidateStruct(payload)
	if errs != nil {
		resultError.Detail = utils.ShowErrors(errs)
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	nanoId, err := gonanoid.New()
	if err != nil {
		resultError.Detail = err
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	dateStart, err := utils.ParseTimeFrStringV2("2006-01-02", payload.DateStart)
	if err != nil {
		resultError.Detail = err
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	dateEnd, err := utils.ParseTimeFrStringV2("2006-01-02", payload.DateEnd)
	if err != nil {
		resultError.Detail = err
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	employeeId, ok := ctx.Locals("user_id").(string)
	if !ok {
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
		Detail: *promotion,
	})

}
