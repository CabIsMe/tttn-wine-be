package delivery

import (
	"net/http"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-module/carbon/v2"
	"go.uber.org/zap"
)

func (h *handlers) Health(ctx *fiber.Ctx) error {
	connection, err := utils.CheckDBConnection(internal.Db)
	if !connection {
		internal.Log.Error("Lỗi kết nối database!", zap.Error(err))
		return ctx.Status(413).JSON(models.RespLocal{
			StatusCode: 413,
			Message:    "DB Connect Failed!",
		})
	}
	return ctx.Status(http.StatusOK).JSON(models.RespLocal{
		StatusCode: 200,
		Message:    "DB Connected!",
	})
}
func (s *handlers) RequireTokenWeb(ctx *fiber.Ctx) error {
	token := string(ctx.Request().Header.Peek("TOKEN"))
	var body interface{}
	ctx.BodyParser(&body)
	internal.Log.Info("RequireWeb", zap.Any("url", ctx.Context().URI()), zap.Any("authen", token), zap.Any("body", body))
	if utils.IsEmpty(token) {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status: internal.SysStatus.TokenRequired.Status,
			Msg:    internal.SysStatus.TokenRequired.Msg,
		})
	}
	clams := jwt.MapClaims{}
	decodeToken, err := jwt.ParseWithClaims(token, clams, func(t *jwt.Token) (interface{}, error) {
		return []byte(internal.Keys.TOKEN_SECRET_KEY_WEB), nil
	})
	if err != nil {
		jwtErr := err.(*jwt.ValidationError).Errors
		if jwtErr == jwt.ValidationErrorExpired {
			return ctx.Status(fiber.StatusBadRequest).JSON(models.Resp{
				Status: internal.SysStatus.TokenExpired.Status,
				Msg:    internal.SysStatus.TokenExpired.Msg,
			})
		}

	}

	if decodeToken == nil || !decodeToken.Valid {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status: internal.SysStatus.InvalidToken.Status,
			Msg:    internal.SysStatus.InvalidToken.Msg,
		})
	}

	ctx.Locals("customer_id", clams["customerId"])
	ctx.Locals("customer_phone", clams["phone"])
	ctx.Locals("app_version", clams["appVersion"])

	internal.Log.Info("RequireWeb", zap.Any("customer_info_fr_jwt", clams))
	return ctx.Next()
}
func (s *handlers) RequireTokenPortal(ctx *fiber.Ctx) error {
	token := string(ctx.Request().Header.Peek("TOKEN"))
	var body interface{}
	ctx.BodyParser(&body)
	internal.Log.Info("RequireTokenPortal", zap.Any("url", ctx.Context().URI()), zap.Any("header", token), zap.Any("body", body))
	if utils.IsEmpty(token) {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status: internal.SysStatus.TokenRequired.Status,
			Msg:    internal.SysStatus.TokenRequired.Msg,
		})
	}
	clams := jwt.MapClaims{}
	decodeToken, err := jwt.ParseWithClaims(token, clams, func(t *jwt.Token) (interface{}, error) {
		return []byte(internal.Keys.PORTAL_FCONNECT_AUTH_KEY), nil
	})
	if err != nil {
		jwtErr := err.(*jwt.ValidationError).Errors
		if jwtErr == jwt.ValidationErrorExpired {
			return ctx.Status(fiber.StatusBadRequest).JSON(models.Resp{
				Status: internal.SysStatus.TokenExpired.Status,
				Msg:    internal.SysStatus.TokenExpired.Msg,
			})
		}
	}

	if decodeToken == nil || !decodeToken.Valid {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Resp{
			Status: internal.SysStatus.InvalidToken.Status,
			Msg:    internal.SysStatus.InvalidToken.Msg,
		})
	}

	ctx.Locals("user_id", clams["user_id"])
	ctx.Locals("tenant_id", clams["tenant_id"])
	ctx.Locals("short_name", clams["short_name"])

	internal.Log.Info("RequireTokenPortal", zap.Any("customer_info_fr_jwt", clams))
	return ctx.Next()
}
func (h *handlers) RequiredTokenLocalWithTenantId(ctx *fiber.Ctx) error {
	tenantId := string(ctx.Request().Header.Peek("X-Tenant"))
	authToken := string(ctx.Request().Header.Peek("X-Auth-Token"))
	correctToken := utils.GenTenantToken(tenantId, internal.Keys.ServiceKey)
	if authToken != correctToken {
		// Nếu là môi trường staging
		if !internal.Envs.IsProduction {
			return ctx.Status(http.StatusUnauthorized).JSON(models.RespToken{
				Token: correctToken,
			})
		} else {
			return ctx.Status(http.StatusUnauthorized).JSON(internal.SysStatus.InvalidToken)
		}
	}

	return ctx.Next()
}
func (h *handlers) RequiredTokenLocalWithoutTenantId(ctx *fiber.Ctx) error {
	if tenant := string(ctx.Request().Header.Peek("X-Tenant")); tenant != "" {
		authToken := string(ctx.Request().Header.Peek("X-Auth-Token"))
		if clientServiceKey, isExist := internal.Keys.LocalServicesKey[tenant]; isExist {
			correctToken := utils.GenTenantToken(clientServiceKey, internal.Keys.ServiceKey)
			if authToken != correctToken {
				// Nếu là môi trường staging
				if !internal.Envs.IsProduction {
					return ctx.Status(http.StatusUnauthorized).JSON(models.RespToken{
						Token: correctToken,
					})
				}
				return ctx.Status(http.StatusUnauthorized).JSON(internal.SysStatus.InvalidToken)
			}
			ctx.Locals("start_time", carbon.Now().TimestampMilli())
			ctx.Locals("token", authToken)
			return ctx.Next()
		}
	}
	return ctx.Status(http.StatusUnauthorized).JSON(internal.SysStatus.InvalidToken)
}
