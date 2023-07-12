package delivery

import (
	"net/http"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/services"
	"github.com/CabIsMe/tttn-wine-be/internal/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationHandler interface {
	SignUpUserHandler(c *fiber.Ctx) error
	UserLoginHandler(ctx *fiber.Ctx) error
}
type auth_handler struct {
	services.MainServices
}

func NewAuthenticationHandler(s services.MainServices) AuthenticationHandler {
	return &auth_handler{
		s,
	}
}
func (h *auth_handler) SignUpUserHandler(ctx *fiber.Ctx) error {
	var payload models.Account
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(http.StatusOK).JSON(models.Resp{
			Status: internal.CODE_WRONG_PARAMS,
			Msg:    internal.MSG_WRONG_PARAMS,
		})
	}
	internal.Log.Info("SignUpUserHandler", zap.Any("Input", payload))
	// validate required per field
	errs := utils.ValidateStruct(payload)
	if errs != nil {
		return ctx.Status(http.StatusOK).JSON(models.Resp{
			Status: internal.CODE_WRONG_PARAMS,
			Msg:    internal.MSG_WRONG_PARAMS,
			Detail: utils.ShowErrors(errs),
		})

	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(http.StatusOK).JSON(models.Resp{
			Status: internal.CODE_SYSTEM_ERROR,
			Msg:    internal.MSG_SYSTEM_ERROR,
		})
	}
	payload.UserPassword = string(hashedPassword)
	payload.RoleId = 2
	rs := h.MainServices.SignUpUserService(payload)
	if rs != nil {
		return ctx.Status(http.StatusOK).JSON(rs)
	}
	internal.Log.Info("SignUpUser", zap.Any("Input", payload), zap.Any("Output", payload))
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: payload,
	})

}

func (h *auth_handler) UserLoginHandler(ctx *fiber.Ctx) error {
	var body models.Account
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusOK).JSON(models.Resp{
			Status: internal.CODE_WRONG_PARAMS,
			Msg:    internal.MSG_WRONG_PARAMS,
		})
	}
	internal.Log.Info("UserLoginHandler", zap.Any("username", body))
	errs := utils.ValidateStruct(body)
	if errs != nil {
		return ctx.Status(http.StatusOK).JSON(models.Resp{
			Status: internal.CODE_WRONG_PARAMS,
			Msg:    internal.MSG_WRONG_PARAMS,
			Detail: utils.ShowErrors(errs),
		})
	}
	accessToken, refreshToken, err := h.MainServices.AuthenticationService.UserLoginService(body)
	if err != nil {
		return ctx.Status(http.StatusOK).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: map[string]interface{}{
			"token":         accessToken,
			"refresh_token": refreshToken,
		},
	})
}
