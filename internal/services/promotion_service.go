package services

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
	"go.uber.org/zap"
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

func (s *promotion_service) CreatePromotionService(model models.Promotion) *internal.SystemStatus {
	errResult := internal.SystemStatus{
		Status: internal.CODE_DB_FAILED,
		Msg:    internal.MSG_DB_FAILED,
	}
	_, err := s.rp.GetEmployee(model.EmployeeId)
	if err != nil {
		internal.Log.Error("CreatePromotionService", zap.Any("GetEmployee", err.Error()))
		return &errResult
	}
	errCreate := s.rp.CreatePromotion(model)
	if errCreate != nil {
		internal.Log.Error("CreatePromotionService", zap.Any("CreatePromotion", errCreate.Error()))
		return &errResult
	}
	return nil
}
