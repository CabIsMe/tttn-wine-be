package delivery

import (
	"fmt"
	"net/http"
	"os"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/services"
	"github.com/CabIsMe/tttn-wine-be/internal/utils"
	"github.com/gofiber/fiber/v2"
	paypal "github.com/logpacker/PayPal-Go-SDK"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
)

// CLIENT_ID : Abx3-C9VHhLKmQPDxgYdnRV2WuoD_qabH0PojrQf5kv71GLi0uEcu6G4axzIGE5TL8oD5ZUx949A5IoR
type CustomerOrderHandler interface {
	AllCustomerOrdersHandler(ctx *fiber.Ctx) error
	CreateCustomerOrder(ctx *fiber.Ctx) error
	AddProductsToCartHandler(ctx *fiber.Ctx) error
	RemoveProductsToCartHandler(ctx *fiber.Ctx) error
	AllProductsInCartHandler(ctx *fiber.Ctx) error
	UpdateCustomerOrderHandler(ctx *fiber.Ctx) error
	ResultPayment(ctx *fiber.Ctx) error
	GetPaymentById(ctx *fiber.Ctx) error
	PaymentSuccess(ctx *fiber.Ctx) error
}
type c_order_handler struct {
	services.MainServices
}

func NewCustomerOrderHandler(s services.MainServices) CustomerOrderHandler {
	return &c_order_handler{
		s,
	}
}

func (h *c_order_handler) PaymentSuccess(c *fiber.Ctx) error {
	paymentId := c.Query("paymentId")
	payerId := c.Query("PayerID")
	token := c.Query("token")
	fmt.Println(paymentId)
	fmt.Println(payerId)
	fmt.Println(token)
	return c.SendString("success")
}

func (h *c_order_handler) GetPaymentById(ctx *fiber.Ctx) error {
	c, _ := paypal.NewClient("Abx3-C9VHhLKmQPDxgYdnRV2WuoD_qabH0PojrQf5kv71GLi0uEcu6G4axzIGE5TL8oD5ZUx949A5IoR",
		"EAemdpvcm6-ve8EaKo7v67BiRZ5rfVCxcSj0Gj5HVE5CEjU2-b2ZE9_RIaDgukntybIUiNfjKiox8-Ce", paypal.APIBaseSandBox)
	c.SetLog(os.Stdout) // Set log to terminal stdout

	// accessToken, err := c.GetAccessToken()
	order, _ := c.GetOrder("PAYID-MS6UYLY5B214533YG293490N")
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: order,
	})
}

func (h *c_order_handler) ResultPayment(ctx *fiber.Ctx) error {
	resultError := models.Resp{
		Status: internal.CODE_WRONG_PARAMS,
		Msg:    internal.MSG_WRONG_PARAMS,
	}
	var body interface{}
	ctx.BodyParser(&body)
	uri := string(ctx.Request().URI().RequestURI())
	tokenAuth := string(ctx.Request().Header.Peek("token"))
	defer func() {
		internal.Log.Info("ResultPayment", zap.Any("uri", uri), zap.Any("auth", tokenAuth), zap.Any("body", body))
	}()
	var requestBody map[string]string
	if err := ctx.BodyParser(&requestBody); err != nil {
		return ctx.Status(200).JSON(resultError)
	}
	orderId := requestBody["order_id"]
	err := h.MainServices.CustomerOrderService.UpdatePaymentStatusCustomerOrderService(orderId)
	if err != nil {
		return ctx.Status(200).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
	})
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
		PaymentStatus           int8                          `json:"payment_status" validate:"required"`
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
	customerOrderId, _ := gonanoid.New()

	inputData := models.CustomerOrder{
		CustomerOrderId: customerOrderId,
		FullName:        payload.FullName,
		TCreate:         utils.GetTimeUTC7(),
		Address:         payload.Address,
		PhoneNumber:     payload.PhoneNumber,
		CustomerId:      customerId,
		PaymentStatus:   payload.PaymentStatus,
		TDelivery:       nil,
	}

	internal.Log.Info("CreateCustomerOrder", zap.Any("InputData", inputData))
	errCreate := h.MainServices.CustomerOrderService.CreateCustomerOrderService(inputData, payload.CustomerOrderDetailInfo)
	if errCreate != nil {
		return ctx.Status(http.StatusOK).JSON(errCreate)
	}
	output := make(map[string]interface{})
	output["customerOrder"] = inputData
	output["customerOrderDetail"] = payload.CustomerOrderDetailInfo
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: output,
	})
}

func (h *c_order_handler) AddProductsToCartHandler(ctx *fiber.Ctx) error {
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
		internal.Log.Info("AddProductsToCartHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth), zap.Any("body", body))
	}()
	var payload models.Cart
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
	// customer_id from jwt
	payload.CustomerId = customerId
	errAdd := h.MainServices.CustomerOrderService.AddProductsToCartService(payload)
	if errAdd != nil {
		return ctx.Status(http.StatusOK).JSON(errAdd)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
	})
}
func (h *c_order_handler) RemoveProductsToCartHandler(ctx *fiber.Ctx) error {
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
		internal.Log.Info("RemoveProductsToCartHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth), zap.Any("body", body))
	}()
	var payload models.Cart
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
	payload.CustomerId = customerId
	errAdd := h.MainServices.CustomerOrderService.RemoveProductsToCartService(payload)
	if errAdd != nil {
		return ctx.Status(http.StatusOK).JSON(errAdd)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
	})
}
func (h *c_order_handler) AllProductsInCartHandler(ctx *fiber.Ctx) error {
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
		internal.Log.Info("AllProductsInCartService", zap.Any("uri", uri), zap.Any("auth", tokenAuth))
	}()
	listData, err := h.MainServices.CustomerOrderService.AllProductsInCartService(customerId)
	if err != nil {
		internal.Log.Error("AllProductsInCartService", zap.Any("Error", err))
		return ctx.Status(http.StatusOK).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: listData,
	})
}
func (h *c_order_handler) UpdateCustomerOrderHandler(ctx *fiber.Ctx) error {
	resultError := models.Resp{
		Status: internal.CODE_WRONG_PARAMS,
		Msg:    internal.MSG_WRONG_PARAMS,
	}
	var body interface{}
	ctx.BodyParser(&body)
	uri := string(ctx.Request().URI().RequestURI())
	tokenAuth := string(ctx.Request().Header.Peek("Authorization"))
	employee_id, ok := ctx.Locals("user_id").(string)
	if !ok {
		return ctx.Status(http.StatusOK).JSON(resultError)
	}
	defer func() {
		internal.Log.Info("UpdateCustomerOrderHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth), zap.Any("body", body))
	}()
	payload := &struct {
		CustomerOrderId string `json:"customer_order_id" validate:"required"`
		TDelivery       string `json:"t_delivery" validate:"required"`
		DelivererId     string `json:"deliverer_id" validate:"required"`
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
	dateDelivery, errParse := utils.ParseTimeFrStringV2("Y-M-D", payload.TDelivery)
	if errParse != nil {
		internal.Log.Error("dateDelivery", zap.Any("Error", errParse.Error()))
		resultError.Detail = errParse.Error()
		return ctx.Status(http.StatusOK).JSON(resultError)
	}

	inputData := models.UpdatingCustomerOrder{
		CustomerOrderId: payload.CustomerOrderId,
		TDelivery:       dateDelivery,
		Status:          models.Cos.ORDER_CONFIRM.StatusCode,
		EmployeeId:      employee_id,
		DelivererId:     payload.DelivererId,
	}
	errUpdate := h.MainServices.CustomerOrderService.UpdateCustomerOrderService(inputData)
	if errUpdate != nil {
		return ctx.Status(http.StatusOK).JSON(errUpdate)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: payload,
	})

}

func (h *c_order_handler) AllCustomerOrdersHandler(ctx *fiber.Ctx) error {
	uri := string(ctx.Request().URI().RequestURI())
	tokenAuth := string(ctx.Request().Header.Peek("Authorization"))
	defer func() {
		internal.Log.Info("AllCustomerOrdersHandler", zap.Any("uri", uri), zap.Any("auth", tokenAuth))
	}()
	listData, err := h.MainServices.CustomerOrderService.AllCustomerOrdersService()
	if err != nil {
		internal.Log.Error("AllCustomerOrdersHandler", zap.Any("Error", err))
		return ctx.Status(http.StatusOK).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: listData,
	})
}
