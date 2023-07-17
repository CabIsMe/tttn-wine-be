package repositories

import (
	"fmt"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerOrderRepository interface {
	CreateCustomerOrder(customerOrder models.CustomerOrder, listDetails []*models.CustomerOrderDetail) error
	AddProductsToCart(cart models.Cart) error
	RemoveProductsToCart(cart models.Cart) error
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
	err2 := tx.Debug().Model(models.CustomerOrderDetail{}).Create(listDetails)
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
			gorm.Expr(fmt.Sprintf("%s - ?", product.ColumnInventoryNumber()), customerOD.Amount))
		if err3.Error != nil {
			internal.Log.Error("Update Product", zap.Any("Error", err3.Error))
			tx.Rollback()
			return err3.Error
		}
		if err3.RowsAffected < 1 {
			tx.Rollback()
			return gorm.ErrInvalidTransaction
		}
	}
	internal.Log.Info("CreateCustomerOrder", zap.Any("Number of record", err2.RowsAffected))
	tx.Commit()
	return nil
}

func (r *c_order_repos) AddProductsToCart(cart models.Cart) error {
	return internal.Db.Debug().Create(&cart).Error
}
func (r *c_order_repos) RemoveProductsToCart(cart models.Cart) error {
	return internal.Db.Debug().Where("customer_id = ? AND product_id = ?", cart.CustomerId, cart.ProductId).Delete(&cart).Error
}
