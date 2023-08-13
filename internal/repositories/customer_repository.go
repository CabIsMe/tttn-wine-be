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
	UpdateCustomer(customer models.Customer) error
	
}

type customer_repos struct {
}

func NewCustomerRepository() CustomerRepository {
	return &customer_repos{}
}

func (r *customer_repos) GetCustomer(customerId string) (*models.Customer, error) {
	model := &models.Customer{}
	result := internal.Db.Where(fmt.Sprintf("%s = ?", model.ColumnCustomerId()), customerId).First(&model).Error
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
func (r *customer_repos) UpdateCustomer(customer models.Customer) error {
	err := internal.Db.Debug().Model(&models.Customer{}).
		Where(fmt.Sprintf("%s = ?", customer.ColumnCustomerId()), customer.CustomerId).
		Updates(customer)
	if err.Error != nil {
		return err.Error
	}
	internal.Log.Info("UpdateCustomer", zap.Any("Number of records", err.RowsAffected))
	return nil
}
