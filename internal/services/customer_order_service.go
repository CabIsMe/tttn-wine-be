package services

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
)

type CustomerOrderService interface {
	CreateCustomerOrderService(customerOrder models.CustomerOrder, customerOrderDetails []*models.CustomerOrderDetail) *internal.SystemStatus
	AddProductsToCartService(cart models.Cart) *internal.SystemStatus
	RemoveProductsToCartService(cart models.Cart) *internal.SystemStatus
}
type c_order_service struct {
	rp repositories.Repos
}

func NewCustomerOrderService(rp repositories.Repos) CustomerOrderService {
	return &c_order_service{
		rp,
	}
}

func (s *c_order_service) CheckCustomerOrderDetailService(customerOrderDetail models.CustomerOrderDetail) *internal.SystemStatus {
	errResult := internal.SystemStatus{
		Status: internal.CODE_DB_FAILED,
		Msg:    internal.MSG_DB_FAILED,
	}
	product, err := s.rp.GetProduct(customerOrderDetail.ProductId)
	if err != nil {
		return &errResult
	}
	if product == nil {
		errResult.Detail = "Product not found"
		return &errResult
	}
	if product.InventoryNumber < customerOrderDetail.Amount {
		errResult.Detail = "Currently out of stock"
		return &errResult
	}
	return nil
}

func (s *c_order_service) CreateCustomerOrderService(
	customerOrder models.CustomerOrder, customerOrderDetails []*models.CustomerOrderDetail) *internal.SystemStatus {
	errResult := internal.SystemStatus{
		Status: internal.CODE_DB_FAILED,
		Msg:    internal.MSG_DB_FAILED,
	}
	// check employee, customer are exist
	customer, err := s.rp.GetCustomer(customerOrder.CustomerId)
	if err != nil {
		return &errResult
	}
	if customer == nil {
		errResult.Detail = "Customer not found"
		return &errResult
	}
	customerOrderId, _ := gonanoid.New()

	customerOrder.CustomerOrderId = customerOrderId
	customerOrder.Status = models.Cos.UNAPPROVED.StatusCode
	flag := -1
	var inputCustomerOrders []*models.CustomerOrderDetail
	for i := 0; i < len(customerOrderDetails); i++ {
		if i == flag {
			continue
		}
		s.CheckCustomerOrderDetailService(*customerOrderDetails[i])
		for j := i + 1; j < len(customerOrderDetails); j++ {
			if customerOrderDetails[i].ProductId == customerOrderDetails[j].ProductId {
				customerOrderDetails[i].Amount += customerOrderDetails[j].Amount
				customerOrderDetails[i].Cost += customerOrderDetails[j].Cost
				flag = j
			}
		}
		nanoId, _ := gonanoid.New()
		customerOrderDetails[i].CustomerOrderDetailId = nanoId
		customerOrderDetails[i].CustomerOrderId = customerOrder.CustomerOrderId
		inputCustomerOrders = append(inputCustomerOrders, customerOrderDetails[i])
	}
	internal.Log.Info("Record ->", zap.Any("inputCustomerOrders", len(inputCustomerOrders)))
	errCreate := s.rp.CreateCustomerOrder(customerOrder, inputCustomerOrders)
	if errCreate != nil {
		errResult.Detail = errCreate.Error()
		return &errResult
	}
	return nil
}

func (s *c_order_service) AddProductsToCartService(cart models.Cart) *internal.SystemStatus {
	errCreate := s.rp.AddProductsToCart(cart)
	if errCreate != nil {
		return &internal.SystemStatus{
			Status: internal.CODE_DB_FAILED,
			Msg:    internal.MSG_DB_FAILED,
			Detail: errCreate.Error(),
		}
	}
	return nil
}
func (s *c_order_service) RemoveProductsToCartService(cart models.Cart) *internal.SystemStatus {
	errCreate := s.rp.RemoveProductsToCart(cart)
	if errCreate != nil {
		return &internal.SystemStatus{
			Status: internal.CODE_DB_FAILED,
			Msg:    internal.MSG_DB_FAILED,
			Detail: errCreate.Error(),
		}
	}
	return nil
}
