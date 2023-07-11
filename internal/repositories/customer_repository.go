package repositories

import (
	"errors"
	"fmt"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetCustomer(empId string) (*models.Customer, error)
	GetCustomerByEmail(email string) (*models.Customer, error)
}

type customer_repos struct {
}

func NewCustomerRepository() CustomerRepository {
	return &customer_repos{}
}

func (r *customer_repos) GetCustomer(empId string) (*models.Customer, error) {
	model := &models.Customer{}
	result := internal.Db.Where(fmt.Sprintf("%s = ?", model.ColumnCustomerId()), empId).First(&model).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		internal.Log.Error("GetCustomer", zap.Any("Error query", result))
		return nil, nil
	}
	return model, result
}
func (r *customer_repos) GetCustomerByEmail(email string) (*models.Customer, error) {
	model := &models.Customer{}
	result := internal.Db.Where(fmt.Sprintf("%s = ?", model.ColumnEmail()), email).First(&model).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		internal.Log.Error("GetCustomerByEmail", zap.Any("Error query", result))
		return nil, nil
	}
	return model, result
}
