package delivery

import (
	"fc_optimal_assignment/internal"
	"fc_optimal_assignment/internal/models"
	"fc_optimal_assignment/internal/repositories"
	"fc_optimal_assignment/internal/services"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-module/carbon/v2"
	"go.uber.org/zap"
)

type Handlers interface {
	Health(*fiber.Ctx) error
	RequiredTokenLocalWithTenantId(ctx *fiber.Ctx) error
	RequiredTokenLocalWithoutTenantId(ctx *fiber.Ctx) error
	RequireTokenPortal(ctx *fiber.Ctx) error
	RequireTokenWeb(ctx *fiber.Ctx) error
	RequiredChecksum(ctx *fiber.Ctx) error
	IsValidTenant(ctx *fiber.Ctx) error
	//
	ScanReadyTasks(ctx *fiber.Ctx) error
	UploadFile(ctx *fiber.Ctx) error
	AutoAssignTasks(ctx *fiber.Ctx) error
	AutoAssignSeller(ctx *fiber.Ctx) error
	//webkit handler
	WebHandler
	/*
			post middleware:

		required ctx.Locals keys: start_time(mili), token, func_name, output,
	*/
	Send2Kibana(ctx *fiber.Ctx) error

	/*
			post middleware:

		required ctx.Locals keys: start_time(mili), func_name

		optional ctx.Locals keys: action_name, app_version, phone, customer_id, status, msg
	*/
	Send2KibanaAll(ctx *fiber.Ctx) error
	// Import Another Handler
	LocationHandler
	ReportHandler
	// Sku
	SkuCoefficientHandler
	StaffByLeaderHandler
}
type handlers struct {
	repositories.Repos
	services.FcServices
	LocationHandler
	ReportHandler
	SkuCoefficientHandler
	StaffByLeaderHandler
	WebHandler
}

func NewHandlers(
	rp repositories.Repos,
	services services.FcServices,
) Handlers {
	return &handlers{
		rp,
		services,
		NewLocationHandler(rp, services),
		NewReportHandler(rp, services),
		NewSkuCoefficientHandler(services),
		NewStaffByLeaderHandler(services),
		NewWebHandler(services),
	}
}
func (h *handlers) ScanReadyTasks(ctx *fiber.Ctx) error {
	ctx.Locals("start_time", carbon.Now().TimestampMilli())
	ctx.Locals("func_name", "ScanReadyTasks")
	h.FcServices.ScanReadyTasksToAssignTenant()
	return ctx.Status(http.StatusOK).JSON(nil)
}
func (h *handlers) AutoAssignSeller(ctx *fiber.Ctx) error {
	ctx.Locals("start_time", carbon.Now().TimestampMilli())
	ctx.Locals("func_name", "ScanReadyTasks")
	rAutoAssignSeller := &models.RAssignSeller{}
	if errParser := ctx.BodyParser(rAutoAssignSeller); errParser != nil {
		internal.Log.Error("BodyParser", zap.Error(errParser))
		return ctx.Status(http.StatusBadRequest).JSON(internal.SysStatus.WrongParams)
	}
	if resp, err := h.SellerAssignment.AutoAssignSeller(rAutoAssignSeller); err != nil {
		return ctx.Status(http.StatusOK).JSON(models.Resp{
			Status: 0,
			Msg:    internal.MSG_SYSTEM_ERROR,
			Detail: err.Error(),
		})
	} else {
		ctx.Locals("status", 1)
		ctx.Locals("msg", "Ok")
		ctx.Locals("output", models.Resp{
			Status: 1,
			Msg:    "OK",
			Detail: resp,
		})

		return ctx.Next()
	}
}
func (h *handlers) AutoAssignTasks(ctx *fiber.Ctx) error {
	internal.Log.Info("Starting AutoAssignTasks ...")
	ctx.Locals("func_name", "AutoAssignTasks")
	ctx.Locals("action_name", "AutoAssignTasks")
	rAutoAssignTask := &models.RAutoAssignTasks{}
	if errParser := ctx.BodyParser(rAutoAssignTask); errParser != nil {
		internal.Log.Error("BodyParser", zap.Error(errParser))
		return ctx.Status(http.StatusBadRequest).JSON(internal.SysStatus.WrongParams)
	}
	if rAutoAssignTask.Tasks != nil {
		// Gắn mặc định
		for _, task := range rAutoAssignTask.Tasks {
			// Mặc định task sẽ được phân công làm sớm - nếu khách hàng đồng ý đặt hẹn lại sớm hơn dự kiến
			isAllowEarlyServe := true
			if task != nil && task.IsEarlyServe == nil {
				task.IsEarlyServe = &isAllowEarlyServe
			}
		}
	}
	internal.Log.Info("", zap.Any("rAutoAssignTask", rAutoAssignTask))
	if resp, err := h.FcServices.AutoAssignTasks(ctx.Context(), rAutoAssignTask); err != nil {
		return ctx.Status(http.StatusOK).JSON(models.Resp{
			Status: 0,
			Msg:    internal.MSG_SYSTEM_ERROR,
			Detail: err.Error(),
		})
	} else {
		ctx.Locals("status", 1)
		ctx.Locals("msg", "Ok")
		ctx.Locals("output", models.Resp{
			Status: 1,
			Msg:    "OK",
			Detail: resp,
		})
		return ctx.Next()
	}
}

func (h *handlers) UploadFile(ctx *fiber.Ctx) error {

	bodyData := &struct {
		SubFolder    string `json:"sub_folder"`
		DataContent  string `json:"data_content"`
		DataType     string `json:"data_type"`
		DataFileName string `json:"data_file_Name"`
		BigFile      int    `json:"big_file"`
	}{}
	err := ctx.BodyParser(&bodyData)
	if err != nil {
		internal.Log.Debug("Cannot parse", zap.Any("body", string(ctx.Body())), zap.Error(err))
		return ctx.Status(fiber.StatusOK).JSON(internal.SysStatus.WrongParams)
	}
	result, errService := h.FcServices.MinioUploadFile(bodyData.SubFolder, bodyData.DataContent, bodyData.DataType, bodyData.DataFileName, bodyData.BigFile)
	if errService != nil {
		if result != nil {
			errService.Msg = fmt.Sprintf("%v", result)
		}
		return ctx.Status(http.StatusOK).JSON(errService)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: result,
	})
}
