package repositories

import (
	"fmt"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"go.uber.org/zap"
)

type CustomerOrderRepository interface {
	CreateCustomerOrder(customerOrder models.CustomerOrder, listDetails []*models.CustomerOrderDetail) error
}

type c_order_repos struct {
}

func NewCustomerOrderRepository() CustomerOrderRepository {
	return &c_order_repos{}
}
func (r *c_order_repos) CreateCustomerOrder(customerOrder models.CustomerOrder,
	listDetails []*models.CustomerOrderDetail) error {

	tx := internal.Db.Begin()

	err := tx.Debug().Create(&customerOrder).Error
	if err != nil {
		internal.Log.Error("Create CustomerOrder", zap.Any("Error", err.Error()))
		return err
	}
	err2 := tx.Debug().Create(listDetails)
	if err2.Error != nil {
		internal.Log.Error("Create CustomerOrderDetail", zap.Any("Error", err2.Error))
		tx.Rollback()
		return err2.Error
	}
	product := models.Product{}
	for _, customerOD := range listDetails {
		err3 := tx.Debug().Where(fmt.Sprintf("%s = ? AND %s >= %d",
			product.ColumnProductId(), product.ColumnInventoryNumber(), customerOD.Amount), customerOD.ProductId).
			Model(&product).Update(product.ColumnInventoryNumber(),
			fmt.Sprintf("%s - %d", product.ColumnInventoryNumber(), customerOD.Amount))
		if err3.Error != nil {
			internal.Log.Error("Update Product", zap.Any("Error", err3.Error))
			tx.Rollback()
			return err3.Error
		}
	}
	internal.Log.Info("CreateCustomerOrder", zap.Any("Number of record", err2.RowsAffected))
	tx.Commit()
	return nil
}

// func (r *c_order_repos) CreateCustomerOrderDetail(model models.CustomerOrder) error {
// 	return internal.Db.Debug().Create(&model).Error
// }
