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

type CustomerHandler interface {
	GetCustomerInfoHandler(ctx *fiber.Ctx) error
	UpdateCustomerHandler(ctx *fiber.Ctx) error
}
type customer_handler struct {
	services.MainServices
}

func NewCustomerHandler(s services.MainServices) CustomerHandler {
	return &customer_handler{
		s,
	}
}

func (h *customer_handler) GetCustomerInfoHandler(ctx *fiber.Ctx) error {
	resultError := models.Resp{
		Status: internal.CODE_WRONG_PARAMS,
		Msg:    internal.MSG_WRONG_PARAMS,
	}

	uri := string(ctx.Request().URI().RequestURI())
	tokenAuth := string(ctx.Request().Header.Peek("token"))
	customerId, ok := ctx.Locals("user_id").(string)
	if !ok {
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	defer func() {
		internal.Log.Info("RemoveProductsToCartHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth))
	}()

	customer, err := h.MainServices.CustomerService.GetCustomerInfoService(customerId)
	if err != nil {
		return ctx.Status(http.StatusOK).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: customer,
	})
}
func (h *customer_handler) UpdateCustomerHandler(ctx *fiber.Ctx) error {
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
		internal.Log.Info("CreatePromotionHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth), zap.Any("body", body))
	}()
	payload := struct {
		FullName    string `json:"full_name" validate:"required"`
		Gender      int8   `json:"gender"`
		DateOfBirth string `json:"date_of_birth"`
		Address     string `json:"address" validate:"required"`
		PhoneNumber string `json:"phone_number" validate:"max=11,min=10"`
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
	dateOfBirth, errParse := utils.ParseTimeFrStringV2("Y-M-D", payload.DateOfBirth)
	if errParse != nil {
		internal.Log.Error("dateOfBirth", zap.Any("Error", errParse.Error()))
		resultError.Detail = errParse.Error()
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	customerInput := models.Customer{
		CustomerId:  customerId,
		FullName:    payload.FullName,
		Gender:      payload.Gender,
		DateOfBirth: dateOfBirth,
		Address:     payload.Address,
		PhoneNumber: payload.PhoneNumber,
	}
	errUpdate := h.MainServices.CustomerService.UpdateCustomerService(customerInput)
	if errUpdate != nil {
		return ctx.Status(http.StatusOK).JSON(errUpdate)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: payload,
	})
}
