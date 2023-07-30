package services

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
)

type AccountService interface {
	GetAccountInfoService(user_id string) interface{}
}
type account_service struct {
	rp repositories.Repos
}

func NewAccountService(rp repositories.Repos) AccountService {
	return &account_service{
		rp,
	}
}
func (s *account_service) GetAccountInfoService(user_id string) interface{} {
	model, err := s.rp.GetEmployee(user_id)
	if err != nil {
		return internal.SysStatus.DbFailed
	}
	result := make(map[string]interface{})
	if model == nil {
		internal.Log.Error("Employee not found")
		return result
	}
	account, err := s.rp.GetAccountInfo(model.Email)
	if err != nil {
		internal.Log.Error("Account not found")
		return result
	}
	if account == nil {
		return result
	}
	result["result"] = account
	return models.Resp{
		Status: 1,
		Msg:    "OK",
		Detail: result,
	}
}
