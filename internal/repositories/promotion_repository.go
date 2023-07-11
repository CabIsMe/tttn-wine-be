package repositories

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
)

type PromotionRepository interface {
	CreatePromotion(model models.Promotion) error
}

type promotion_repos struct {
}

func NewPromotionRepository() PromotionRepository {
	return &promotion_repos{}
}

func (r *promotion_repos) CreatePromotion(model models.Promotion) error {
	return internal.Db.Debug().Create(model).Error
}
