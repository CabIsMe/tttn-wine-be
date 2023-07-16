package repositories

import (
	"errors"
	"fmt"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateAccountUser(payload models.Account) error
	GetAccountByUsername(username string) (*models.Account, error)
}

type auth_repos struct {
}

func NewAuthRepository() AuthRepository {
	return &auth_repos{}
}

func (r *auth_repos) CreateAccountUser(payload models.Account) error {
	return internal.Db.Debug().Create(&payload).Error
}

func (r *auth_repos) GetAccountByUsername(username string) (*models.Account, error) {
	model := &models.Account{}
	result := internal.Db.Where(fmt.Sprintf("%s = ?", model.ColumnUsername()), username).First(&model).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		internal.Log.Error("GetAccount", zap.Any("Error query", result))
		return nil, nil
	}
	return model, result
}
