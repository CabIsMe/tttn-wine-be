package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerOrderService interface {
	CreateCustomerOrderService(customerOrder models.CustomerOrder, customerOrderDetails []*models.CustomerOrderDetail) *internal.SystemStatus
	UpdateCustomerOrderService(customerOrder models.UpdatingCustomerOrder) *internal.SystemStatus
	AddProductsToCartService(cart models.Cart) *internal.SystemStatus
	RemoveProductsToCartService(cart models.Cart) *internal.SystemStatus
	AllProductsInCartService(customerId string) ([]*models.Cart, *internal.SystemStatus)
	UpdatePaymentStatusCustomerOrderService(customerOrderId string) *internal.SystemStatus
	UpdateStatusCustomerOrderService(customerOrderId string, stt int8) *internal.SystemStatus
	AllCustomerOrdersService() (interface{}, *internal.SystemStatus)
	AllCustomerOrdersByStatusService(listStt []int8) (interface{}, *internal.SystemStatus)
	AllDelivererIdsService() ([]models.Employee, *internal.SystemStatus)
	GetCustomerOrderService(customerOrderId string) (*models.CustomerOrder, *internal.SystemStatus)
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

	customerOrder.Status = models.Cos.UNAPPROVED.StatusCode
	// customerOrder.TCreate = time.g
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
		if errors.Is(errCreate, gorm.ErrInvalidTransaction) {
			errResult.Detail = "Sorry, the product you bought is out of stock"
		}
		return &errResult
	}
	return nil
}

func (s *c_order_service) AddProductsToCartService(cart models.Cart) *internal.SystemStatus {

	errCreate := s.rp.AddProductsToCart(cart)
	if errCreate != nil {
		if strings.Contains(errCreate.Error(), "Duplicate entry") {
			return &internal.SystemStatus{
				Status: internal.CODE_DB_FAILED,
				Msg:    internal.MSG_DB_FAILED,
				Detail: "The product has been added to cart",
			}
		}
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

func (s *c_order_service) AllProductsInCartService(customerId string) ([]*models.Cart, *internal.SystemStatus) {
	errResult := &internal.SystemStatus{
		Status: internal.CODE_DB_FAILED,
		Msg:    internal.MSG_DB_FAILED,
	}
	// listResults := make(map[string]interface{})

	products, err := s.rp.GetAllProductsInCart(customerId)
	if err != nil {
		errResult.Detail = err.Error()
		return nil, errResult
	}
	prom, err := s.rp.GetPromotionByDate()
	if err != nil {
		errResult.Detail = err.Error()
		return nil, errResult
	}
	if prom == nil {
		for _, model := range products {
			product, errPro := s.rp.GetProduct(model.ProductId)
			if errPro != nil {
				errResult.Detail = errPro.Error()
				return nil, errResult
			}
			if product == nil || product.InventoryNumber < model.Amount {
				errResult.Detail = fmt.Sprintf("%s is no longer available or has sold out.", product.ProductName)
				return nil, errResult
			}
			model.Cost = product.Cost
		}

		return products, nil
	}
	for _, model := range products {
		internal.Log.Info("Info products in cart", zap.Any("product", model))
		promDetail, errProm := s.rp.GetPromotionDetail(model.ProductId, prom.PromotionId)
		if errProm != nil {
			errResult.Detail = errProm.Error()
			return nil, errResult
		}
		product, errPro := s.rp.GetProduct(model.ProductId)
		if errPro != nil {
			errResult.Detail = errPro.Error()
			return nil, errResult
		}
		// promotion detaill is null -> assign cost data =  cost product
		if promDetail == nil {
			if product == nil || product.InventoryNumber < model.Amount {
				errResult.Detail = fmt.Sprintf("%s is no longer available or has sold out.", product.ProductName)
				return nil, errResult
			}
			model.Cost = product.Cost
			continue
		}
		// promotion detail not null -> assign cost data= cost product * (1- %discount)
		if product == nil || product.InventoryNumber < model.Amount {
			errResult.Detail = fmt.Sprintf("%s is no longer available or has sold out.", product.ProductName)
			return nil, errResult
		}
		model.Cost = product.Cost * (1 - promDetail.DiscountPercentage)
	}
	return products, nil
}
func (s *c_order_service) UpdateCustomerOrderService(customerOrder models.UpdatingCustomerOrder) *internal.SystemStatus {
	deliverer, err := s.rp.GetEmployee(customerOrder.DelivererId)
	if deliverer == nil || err != nil {
		return internal.SysStatus.DbFailed
	}
	err = s.rp.UpdateCustomerOrder(customerOrder)
	if err != nil {
		return internal.SysStatus.DbFailed
	}
	return nil
}

func (s *c_order_service) UpdatePaymentStatusCustomerOrderService(customerOrderId string) *internal.SystemStatus {
	err := s.rp.UpdatePaymentStatusCustomerOrder(customerOrderId, 2)
	if err != nil {
		return internal.SysStatus.DbFailed
	}
	return nil
}
func (s *c_order_service) UpdateStatusCustomerOrderService(customerOrderId string, stt int8) *internal.SystemStatus {
	err := s.rp.UpdateStatusCustomerOrder(customerOrderId, stt)
	if err != nil {
		return internal.SysStatus.DbFailed
	}
	return nil
}

func (s *c_order_service) AllCustomerOrdersService() (interface{}, *internal.SystemStatus) {
	listResults, err := s.rp.GetAllCustomerOrders()
	if err != nil {
		internal.Log.Error("AllCustomerOrdersService", zap.Any("err", err.Error()))
		return nil, internal.SysStatus.DbFailed
	}
	internal.Log.Info("AllCustomerOrdersService", zap.Any("listResults", listResults))
	return listResults, nil
}
func (s *c_order_service) AllCustomerOrdersByStatusService(listStt []int8) (interface{}, *internal.SystemStatus) {
	listResults, err := s.rp.GetAllCustomerOrdersByStatus(listStt)
	if err != nil {
		internal.Log.Error("AllCustomerOrdersByStatusService", zap.Any("err", err.Error()))
		return nil, internal.SysStatus.DbFailed
	}
	internal.Log.Info("AllCustomerOrdersByStatusService", zap.Any("listResults", listResults))
	return listResults, nil
}

func (s *c_order_service) AllDelivererIdsService() ([]models.Employee, *internal.SystemStatus) {
	listResults, err := s.rp.GetAllDeliverers()
	if err != nil {
		internal.Log.Error("AllDelivererIdsService", zap.Any("err", err.Error()))
		return nil, internal.SysStatus.DbFailed
	}
	internal.Log.Info("AllDelivererIdsService", zap.Any("listResults", listResults))
	return listResults, nil
}
func (s *c_order_service) GetCustomerOrderService(customerOrderId string) (*models.CustomerOrder, *internal.SystemStatus) {
	result, err := s.rp.GetCustomerOrder(customerOrderId)
	if err != nil {
		return nil, internal.SysStatus.DbFailed
	}
	return result, nil
}
