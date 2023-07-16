package services

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type CustomerOrderService interface {
	CreateCustomerOrderService(customerOrder models.CustomerOrder, customerOrderDetails []*models.CustomerOrderDetail) *internal.SystemStatus
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
	// employee, err := s.rp.GetEmployee(customerOrder.EmployeeId)
	// if err != nil {
	// 	return &errResult
	// }
	// if employee == nil {
	// 	errResult.Detail = "Employee not found"
	// 	return &errResult
	// }
	// deliverer, err := s.rp.GetEmployee(customerOrder.DelivererId)
	// if err != nil {
	// 	return &errResult
	// }
	// if deliverer == nil {
	// 	errResult.Detail = "Deliverer not found"
	// 	return &errResult
	// }
	customerOrderId, _ := gonanoid.New()

	customerOrder.CustomerOrderId = customerOrderId
	customerOrder.Status = models.Cos.UNAPPROVED.StatusCode
	for _, model := range customerOrderDetails {
		s.CheckCustomerOrderDetailService(*model)
		nanoId, _ := gonanoid.New()
		model.CustomerOrderDetailId = nanoId
		model.CustomerOrderId = customerOrder.CustomerOrderId
	}
	errCreate := s.rp.CreateCustomerOrder(customerOrder, customerOrderDetails)
	if errCreate != nil {
		errResult.Detail = errCreate.Error()
		return &errResult
	}
	return nil
}
