package repositories

import (
	"errors"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BillRepository interface {
	GetBillByCustomerOrderId(customerOrderId string) (*models.Bill, error)
	CreateBill(bill models.Bill) error
}

type bill_repos struct {
}

func NewBillRepository() BillRepository {
	return &bill_repos{}
}
func (r *bill_repos) GetBillByCustomerOrderId(customerOrderId string) (*models.Bill, error) {
	var model = &models.Bill{}
	err := internal.Db.Debug().Table(model.TableName()).Select("*").
		Where("customer_order_id = ?", customerOrderId).
		Take(&model)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		internal.Log.Error("GetBillByCustomerOrderId", zap.Any("Error query", err.Error))
		return nil, nil
	}
	if err.Error != nil {
		internal.Log.Error("GetBillByCustomerOrderId", zap.Any("Error", err.Error))
		return nil, err.Error
	}

	return model, nil
}

func (r *bill_repos) CreateBill(bill models.Bill) error {
	return internal.Db.Debug().Create(bill).Error
}
