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

type BillHandler interface {
	GetBillByCustomerOrderIdHandler(ctx *fiber.Ctx) error
	CreateBillHandler(ctx *fiber.Ctx) error
}
type bill_handler struct {
	services.MainServices
}

func NewBillHandler(s services.MainServices) BillHandler {
	return &bill_handler{
		s,
	}
}
func (h *bill_handler) GetBillByCustomerOrderIdHandler(ctx *fiber.Ctx) error {
	resultError := models.Resp{
		Status: internal.CODE_WRONG_PARAMS,
		Msg:    internal.MSG_WRONG_PARAMS,
	}
	var body interface{}
	ctx.BodyParser(&body)
	uri := string(ctx.Request().URI().RequestURI())
	tokenAuth := string(ctx.Request().Header.Peek("token"))
	defer func() {
		internal.Log.Info("GetBillByCustomerOrderIdHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth), zap.Any("body", body))
	}()
	payload := struct {
		CustomerOrderId string `json:"customer_order_id" validate:"required"`
	}{}
	if err := ctx.BodyParser(&payload); err != nil {
		internal.Log.Error("BodyParser", zap.Any("Error", err.Error()))
		resultError.Detail = err.Error()
		return ctx.Status(200).JSON(resultError)
	}
	errs := utils.ValidateStruct(payload)
	if errs != nil {
		internal.Log.Error("ValidateStruct", zap.Any("Error", utils.ShowErrors(errs)))
		resultError.Detail = utils.ShowErrors(errs)
		return ctx.Status(http.StatusOK).JSON(resultError)
	}

	model, errGet := h.MainServices.BillService.GetBillByCustomerOrderIdService(payload.CustomerOrderId)
	if errGet != nil {
		return ctx.Status(http.StatusOK).JSON(errGet)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: model,
	})
}
func (h *bill_handler) CreateBillHandler(ctx *fiber.Ctx) error {
	resultError := models.Resp{
		Status: internal.CODE_WRONG_PARAMS,
		Msg:    internal.MSG_WRONG_PARAMS,
	}
	var body interface{}
	ctx.BodyParser(&body)
	uri := string(ctx.Request().URI().RequestURI())
	tokenAuth := string(ctx.Request().Header.Peek("token"))
	defer func() {
		internal.Log.Info("CreateBillHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth), zap.Any("body", body))
	}()
	payload := &struct {
		TaxId           string `json:"tax_id"`
		TaxName         string `json:"tax_name"`
		CustomerOrderId string `json:"customer_order_id" validate:"required"`
	}{}
	if err := ctx.BodyParser(&payload); err != nil {
		internal.Log.Error("BodyParser", zap.Any("Error", err.Error()))
		resultError.Detail = err.Error()
		return ctx.Status(200).JSON(resultError)
	}
	errs := utils.ValidateStruct(payload)
	if errs != nil {
		internal.Log.Error("ValidateStruct", zap.Any("Error", utils.ShowErrors(errs)))
		resultError.Detail = utils.ShowErrors(errs)
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	employeeId, ok := ctx.Locals("user_id").(string)
	if !ok {
		internal.Log.Error("employeeId, ok", zap.Any("Error", ok))
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	inputData := models.Bill{
		TCreate:         utils.GetTimeUTC7(),
		TaxId:           payload.TaxId,
		TaxName:         payload.TaxName,
		CustomerOrderId: payload.CustomerOrderId,
		EmployeeId:      employeeId,
	}
	errGet := h.MainServices.BillService.CreateBillService(inputData)
	if errGet != nil {
		return ctx.Status(http.StatusOK).JSON(errGet)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
	})
}
