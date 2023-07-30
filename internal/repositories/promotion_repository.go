package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PromotionRepository interface {
	CreatePromotion(model models.Promotion) error
	CheckLogicalPromotionDate(dateInput time.Time) (bool, error)
	GetAllPromotions() ([]models.Promotion, error)
	GetPromotionByDate() (*models.PromotionAndPercent, error)
	GetPromotionDetail(productId string, promotionId string) (*models.PromotionDetail, error)
	CreatePromotionDetail(model models.PromotionDetail) error
}

type promotion_repos struct {
}

func NewPromotionRepository() PromotionRepository {
	return &promotion_repos{}
}

func (r *promotion_repos) CreatePromotion(model models.Promotion) error {
	return internal.Db.Debug().Create(&model).Error
}
func (r *promotion_repos) CheckLogicalPromotionDate(dateInput time.Time) (bool, error) {
	model := &models.Promotion{}
	var exists bool
	err := internal.Db.Debug().Model(model).
		Select("count(*) > 0").
		Where(fmt.Sprintf("%s > ?", model.ColumnDateEnd()), dateInput).
		Find(&exists).Error
	return exists, err
}
func (r *promotion_repos) GetAllPromotions() ([]models.Promotion, error) {
	var listData []models.Promotion
	result := internal.Db.Debug().Model(&models.Promotion{}).Find(&listData)
	if result.Error != nil {
		return nil, result.Error
	}
	internal.Log.Info("GetAllPromotions", zap.Any("Number of records: ", result.RowsAffected))
	return listData, nil
}
func (r *promotion_repos) CreatePromotionDetail(model models.PromotionDetail) error {
	return internal.Db.Debug().Create(&model).Error
}
func (r *promotion_repos) GetPromotionDetail(productId string, promotionId string) (*models.PromotionDetail, error) {
	model := &models.PromotionDetail{}
	result := internal.Db.Where(fmt.Sprintf("%s = ? AND %s = ?", model.ColumnProductId(), model.ColumnPromotionId()), productId, promotionId).
		Take(&model).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		internal.Log.Error("GetPromotionDetail", zap.Any("Error query", result))
		return nil, nil
	}
	return model, result
}

func (r *promotion_repos) GetPromotionByDate() (*models.PromotionAndPercent, error) {
	model := &models.PromotionAndPercent{}
	now := utils.GetTimeUTC7()
	result := internal.Db.Select("promotion.*, promotion_detail.discount_percentage").
		Table("promotion").
		Joins("INNER JOIN promotion_detail ON promotion.promotion_id = promotion_detail.promotion_id").
		Where("promotion.date_start <= ? AND promotion.date_end >= ?", now, now).
		Order("promotion_detail.discount_percentage DESC").
		First(model).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		internal.Log.Error("GetPromotionByDate", zap.Any("Error query", result))
		return nil, nil
	}
	return model, result
}
