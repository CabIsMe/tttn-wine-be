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
	// confirm order, appoint employee for delivery, update time delivery
	UpdateCustomerOrder(customerOrder models.UpdatingCustomerOrder) error
	AddProductsToCart(cart models.Cart) error
	RemoveProductsToCart(cart models.Cart) error
	GetAllProductsInCart(customerId string) ([]*models.Cart, error)
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
		// inventory = 10. order 2 -> update 10-2
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
		// delete in cart
		err4 := tx.Debug().Where("customer_id = ? AND product_id = ?", customerOrder.CustomerId, customerOD.ProductId).Delete(&models.Cart{}).Error
		if err4 != nil {
			internal.Log.Error("Create CustomerOrder", zap.Any("Error", err4))
			tx.Rollback()
			return err4
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
func (r *c_order_repos) GetAllProductsInCart(customerId string) ([]*models.Cart, error) {
	var products []*models.Cart
	err := internal.Db.Debug().Select("*").
		Table("cart").
		Preload("ProductInfo").
		Find(&products)
	if err.Error != nil {
		return nil, err.Error
	}
	internal.Log.Info("GetAllProductsInCart", zap.Any("Number of records: ", err.RowsAffected))
	return products, nil
}

func (r *c_order_repos) UpdateCustomerOrder(customerOrder models.UpdatingCustomerOrder) error {
	fmt.Println(customerOrder)
	result := internal.Db.Debug().Model(&models.CustomerOrder{}).
		Where(fmt.Sprintf("customer_order_id = ?"), customerOrder.CustomerOrderId).
		Updates(&models.CustomerOrder{
			TDelivery:   &customerOrder.TDelivery,
			Status:      customerOrder.Status,
			DelivererId: &customerOrder.DelivererId,
			EmployeeId:  &customerOrder.EmployeeId,
		})
	if result.Error != nil {
		return result.Error
	}
	internal.Log.Info("UpdateCustomerOrder", zap.Any("Number of records", result.RowsAffected))
	return nil
}
