package repositories

import (
	"fmt"
	"time"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"go.uber.org/zap"
)

type PromotionRepository interface {
	CreatePromotion(model models.Promotion) error
	CheckLogicalPromotionDate(dateInput time.Time) (bool, error)
	GetAllPromotions() ([]models.Promotion, error)
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
