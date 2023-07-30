package services

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
	"go.uber.org/zap"
)

type PromotionService interface {
	CreatePromotionService(model models.Promotion) *internal.SystemStatus
	CreatePromotionDetailService(model models.PromotionDetail) *internal.SystemStatus
	GetPromotionByDateService() (*models.PromotionAndPercent, *internal.SystemStatus)
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
	isExist, err := s.rp.CheckLogicalPromotionDate(model.DateStart)
	if err != nil {
		internal.Log.Error("CreatePromotionService", zap.Any("CheckLogicalPromotionDate", err.Error()))
		return &errResult
	}
	if isExist || model.DateEnd.Unix() <= model.DateStart.Unix() {
		return &internal.SystemStatus{
			Status: internal.CODE_WRONG_PARAMS,
			Msg:    internal.MSG_WRONG_PARAMS,
			Detail: "Date input invalid",
		}
	}
	errCreate := s.rp.CreatePromotion(model)
	if errCreate != nil {
		internal.Log.Error("CreatePromotionService", zap.Any("CreatePromotion", errCreate.Error()))
		return &errResult
	}
	return nil
}
func (s *promotion_service) CreatePromotionDetailService(modelInput models.PromotionDetail) *internal.SystemStatus {
	_, err := s.rp.GetPromotionDetail(modelInput.ProductId, modelInput.PromotionId)
	errResult := internal.SystemStatus{
		Status: internal.CODE_DB_FAILED,
		Msg:    internal.MSG_DB_FAILED,
	}
	if err != nil {
		return &errResult
	}
	errCreate := s.rp.CreatePromotionDetail(modelInput)
	if errCreate != nil {
		return &errResult
	}
	return nil
}
func (s *promotion_service) GetPromotionByDateService() (*models.PromotionAndPercent, *internal.SystemStatus) {
	prom, err := s.rp.GetPromotionByDate()
	if err != nil {
		return nil, internal.SysStatus.DbFailed
	}
	return prom, nil
}
