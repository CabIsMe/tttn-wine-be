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
	CreateAccountUser(payload models.AccountAndFullName) error
	GetAccountByUsername(username string) (*models.Account, error)
}

type auth_repos struct {
}

func NewAuthRepository() AuthRepository {
	return &auth_repos{}
}

func (r *auth_repos) CreateAccountUser(payload models.AccountAndFullName) error {
	tx := internal.Db.Begin()
	err := tx.Debug().Model(&models.Account{}).Create(models.Account{
		Username:     payload.Email,
		UserPassword: payload.Password,
		RoleId:       payload.RoleId,
	}).Error
	if err != nil {
		internal.Log.Error("CreateAccountUser", zap.Any("Error query", err))
		tx.Rollback()
		return err
	}
	err = tx.Debug().Table("customer").Select("customer_id", "email", "full_name").Create(&models.Customer{
		CustomerId: payload.UserId,
		FullName:   payload.Name,
		Email:      payload.Email,
	}).Error
	if err != nil {
		internal.Log.Error("CreateAccountUser", zap.Any("Error query", err))
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
	// return internal.Db.Debug().Create(&payload).Error
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
