package services

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
)

type PromotionService interface {
	CreatePromotionService(model models.Promotion) *internal.SystemStatus
}
type promotion_service struct {
	rp repositories.Repos
}

func NewPromotionService(rp repositories.Repos) PromotionService {
	return &promotion_service{
		rp,
	}
}

func (s *product_service) CreatePromotionService(model models.Promotion) *internal.SystemStatus {
	errResult := internal.SystemStatus{
		Status: internal.CODE_DB_FAILED,
		Msg:    internal.MSG_DB_FAILED,
	}
	empl, err := s.rp.GetEmployee(model.EmployeeId)
	if err != nil {
		return &errResult
	}
	if empl == nil {
		errResult.Detail = "Employee not found"
		return &errResult
	}
	err = s.rp.CreatePromotion()
}
