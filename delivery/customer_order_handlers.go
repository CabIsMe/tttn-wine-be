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

type CustomerOrderHandler interface {
	CreateCustomerOrder(ctx *fiber.Ctx) error
}
type c_order_handler struct {
	services.MainServices
}

func NewCustomerOrderHandler(s services.MainServices) CustomerOrderHandler {
	return &c_order_handler{
		s,
	}
}
func (h *c_order_handler) CreateCustomerOrder(ctx *fiber.Ctx) error {
	resultError := models.Resp{
		Status: internal.CODE_WRONG_PARAMS,
		Msg:    internal.MSG_WRONG_PARAMS,
	}
	var body interface{}
	ctx.BodyParser(&body)
	uri := string(ctx.Request().URI().RequestURI())
	tokenAuth := string(ctx.Request().Header.Peek("token"))
	customerId, ok := ctx.Locals("user_id").(string)
	if !ok {
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	defer func() {
		internal.Log.Info("CreateCustomerOrder", zap.Any("uri", uri), zap.Any("auth", tokenAuth), zap.Any("body", body))
	}()

	type customerOrderInput struct {
		FullName                string                        `json:"full_name"`
		Address                 string                        `json:"address" validate:"required"`
		PhoneNumber             string                        `json:"phone_number" validate:"required"`
		TDelivery               string                        `json:"t_delivery" validate:"required"`
		Status                  string                        `json:"status"`
		EmployeeId              string                        `json:"employee_id"`
		DelivererId             string                        `json:"deliverer_id"`
		CustomerOrderDetailInfo []*models.CustomerOrderDetail `json:"customer_order_detail_info"`
	}
	var payload customerOrderInput
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
	payload.Status = models.Cos.UNAPPROVED.StatusDesc
	dateDelivery, errParse := utils.ParseTimeFrStringV2("Y-M-D", payload.TDelivery)
	if errParse != nil {
		internal.Log.Error("dateDelivery", zap.Any("Error", errParse.Error()))
		resultError.Detail = errParse.Error()
		return ctx.Status(http.StatusOK).JSON(resultError)
	}

	inputData := models.CustomerOrder{
		FullName:    payload.FullName,
		TCreate:     utils.GetTimeUTC7(),
		Address:     payload.Address,
		PhoneNumber: payload.PhoneNumber,
		TDelivery:   dateDelivery,
		EmployeeId:  "",
		DelivererId: "",
		CustomerId:  customerId,
	}

	internal.Log.Info("CreateCustomerOrder", zap.Any("InputData", inputData))
	errCreate := h.MainServices.CustomerOrderService.CreateCustomerOrderService(inputData, payload.CustomerOrderDetailInfo)
	if errCreate != nil {
		return ctx.Status(http.StatusOK).JSON(errCreate)
	}
	return nil
}
