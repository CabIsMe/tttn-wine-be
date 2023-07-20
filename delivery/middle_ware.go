package delivery

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

func (h *handlers) VerifyTokenClient(ctx *fiber.Ctx) error {
	token := string(ctx.Request().Header.Peek("token"))
	internal.Log.Info("VerifyTokenClient", zap.Any("header", token))
	if utils.IsEmpty(token) {
		return ctx.Status(fiber.StatusOK).JSON(models.Resp{
			Status: internal.CODE_TOKEN_REQUIRED,
			Msg:    internal.MSG_TOKEN_REQUIRED,
		})
	}
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(internal.Keys.ACCESS_TOKEN_SECRET), nil
	})
	if err != nil {
		jwtErr := err.(*jwt.ValidationError).Errors
		if jwtErr == jwt.ValidationErrorExpired {
			return ctx.Status(fiber.StatusOK).JSON(models.Resp{
				Status: internal.CODE_TOKEN_EXPIRED,
				Msg:    internal.MSG_TOKEN_EXPIRED,
			})
		}
	}
	if parsedToken == nil || !parsedToken.Valid {
		return ctx.Status(fiber.StatusOK).JSON(models.Resp{
			Status: internal.CODE_INVALID_TOKEN,
			Msg:    internal.MSG_INVALID_TOKEN,
		})
	}
	ctx.Locals("user_id", claims["user_id"])
	internal.Log.Info("VerifyTokenClient", zap.Any("info", claims))
	return ctx.Next()
}

func (h *handlers) VerifyTokenInside(ctx *fiber.Ctx) error {
	token := string(ctx.Request().Header.Peek("Authorization"))
	internal.Log.Info("VerifyTokenInside", zap.Any("header", token))
	if utils.IsEmpty(token) {
		return ctx.Status(fiber.StatusOK).JSON(models.Resp{
			Status: internal.CODE_TOKEN_REQUIRED,
			Msg:    internal.MSG_TOKEN_REQUIRED,
		})
	}
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(internal.Keys.INSIDE_ACCESS_TOKEN_SECRET), nil
	})
	if err != nil {
		jwtErr := err.(*jwt.ValidationError).Errors
		if jwtErr == jwt.ValidationErrorExpired {
			return ctx.Status(fiber.StatusOK).JSON(models.Resp{
				Status: internal.CODE_TOKEN_EXPIRED,
				Msg:    internal.MSG_TOKEN_EXPIRED,
			})
		}
	}
	if parsedToken == nil || !parsedToken.Valid {
		return ctx.Status(fiber.StatusOK).JSON(models.Resp{
			Status: internal.CODE_INVALID_TOKEN,
			Msg:    internal.MSG_INVALID_TOKEN,
		})
	}
	ctx.Locals("user_id", claims["user_id"])
	internal.Log.Info("VerifyTokenInside", zap.Any("info", claims))
	return ctx.Next()
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
		return []byte(internal.Keys.ACCESS_TOKEN_SECRET), nil
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

// func (h *handlers) checkRequiredFields(requiredFields []string) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		// Get the parsed body
// 		body := c.Locals("body").(map[string]interface{})

// 		// Check if all required fields are present
// 		for _, field := range requiredFields {
// 			if _, exists := body[field]; !exists {
// 				// If a required field is missing, return a bad request error
// 				return c.Status(http.StatusOK).JSON(models.Resp{
// 					Status: internal.CODE_WRONG_PARAMS,
// 					Msg:    internal.MSG_WRONG_PARAMS,
// 					Detail: "Missing '" + field + "' parameter",
// 				})
// 			}
// 		}

// 		// Proceed to the next handler if all required fields are present
// 		return c.Next()
// 	}
// }
