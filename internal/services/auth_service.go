package services

import (
	"time"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
	"github.com/CabIsMe/tttn-wine-be/internal/utils"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService interface {
	SignUpUserService(model models.Account) *internal.SystemStatus
	UserLoginService(model models.Account) (interface{}, interface{}, *internal.SystemStatus)
}
type auth_service struct {
	rp repositories.Repos
}

func NewAuthenticationService(rp repositories.Repos) AuthenticationService {
	return &auth_service{
		rp,
	}
}
func (s *auth_service) SignUpUserService(model models.Account) *internal.SystemStatus {
	errResult := internal.SystemStatus{
		Status: internal.CODE_DB_FAILED,
		Msg:    internal.MSG_DB_FAILED,
	}

	_, err := s.rp.GetAccountByUsername(model.Username)
	if err != nil {
		internal.Log.Error("Error in SignUpUserService", zap.Any("model", model))
		return &errResult
	}
	err = s.rp.CreateAccountUser(model)
	if err != nil {
		internal.Log.Error("Error in SignUpUserService", zap.Any("model", model))
		errResult.Detail = "Đã tồn tại " + model.Username
		return &errResult
	}
	return nil
}

func (s *auth_service) UserLoginService(model models.Account) (interface{}, interface{}, *internal.SystemStatus) {
	errResult := internal.SystemStatus{
		Status: internal.CODE_DB_FAILED,
		Msg:    internal.MSG_DB_FAILED,
	}
	user, err := s.rp.GetAccountByUsername(model.Username)
	if err != nil {
		internal.Log.Error("Error in UserLoginService", zap.Any("model", model), zap.Error(err))
		return nil, nil, &errResult
	}
	if user == nil {
		internal.Log.Error("Error in UserLoginService", zap.Any("model", model), zap.Any("user", user))
		return nil, nil, &errResult
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(model.UserPassword)); err != nil {
		internal.Log.Error("Error in UserLoginService", zap.Any("model", model), zap.Error(err))
		return nil, nil, &internal.SystemStatus{
			Status: internal.CODE_DB_FAILED,
			Msg:    internal.MSG_DB_FAILED,
			Detail: "Mật khẩu không chính xác",
		}
	}
	var accessToken, refreshToken string
	var err2, err3 error
	if user.RoleId == 2 { // customer
		customer, err := s.rp.GetCustomerByEmail(user.Username)
		if err != nil {
			internal.Log.Error("Error in UserLoginService", zap.Any("model", model), zap.Error(err))
			return nil, nil, &errResult
		}
		if customer == nil {
			return nil, nil, &internal.SystemStatus{
				Status: internal.CODE_WRONG_PARAMS,
				Msg:    internal.MSG_WRONG_PARAMS,
			}
		}
		internal.Log.Info("Login in with Customer")
		accessToken, err2 = generateUserToken(user.Username, customer.CustomerId, user.RoleId, 24, internal.Keys.ACCESS_TOKEN_SECRET)
		refreshToken, err3 = generateUserToken(user.Username, customer.CustomerId, -1, 24*15, internal.Keys.REFRESH_TOKEN_SECRET)
	} else { // admin
		employee, err := s.rp.GetEmployeeByEmail(user.Username)
		if err != nil {
			internal.Log.Error("Error in UserLoginService", zap.Any("model", model), zap.Error(err))
			return nil, nil, &errResult
		}
		if employee == nil {
			return nil, nil, &internal.SystemStatus{
				Status: internal.CODE_WRONG_PARAMS,
				Msg:    internal.MSG_WRONG_PARAMS,
			}
		}
		internal.Log.Info("Login in with Employee")
		accessToken, err2 = generateUserToken(user.Username, employee.EmployeeId, user.RoleId, 24, internal.Keys.INSIDE_ACCESS_TOKEN_SECRET)
		refreshToken, err3 = generateUserToken(user.Username, employee.EmployeeId, -1, 24*15, internal.Keys.INSIDE_REFRESH_TOKEN_SECRET)
	}
	if err2 != nil || err3 != nil {
		internal.Log.Error("Error in UserLoginService", zap.Any("model", model), zap.Error(err2), zap.Error(err3))
		return nil, nil, &errResult
	}
	return accessToken, refreshToken, nil
}
func generateUserToken(username string, userId string, role_id int8, expireHours uint16, keyStr string) (string, error) {
	now := utils.GetTimeUTC7()
	claims := jwt.MapClaims{
		"exp":      now.Add(time.Hour * time.Duration(expireHours)).Unix(),
		"username": username,
		"role_id":  role_id,
		"user_id":  userId,
		"issuedAt": now,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(keyStr))
}
