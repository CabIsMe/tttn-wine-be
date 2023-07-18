package services

import (
	"fmt"
	"strings"

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
	AllProductsInCartService(customerId string) ([]models.Cart, *internal.SystemStatus)
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

func (s *c_order_service) AllProductsInCartService(customerId string) ([]models.Cart, *internal.SystemStatus) {
	errResult := &internal.SystemStatus{
		Status: internal.CODE_DB_FAILED,
		Msg:    internal.MSG_DB_FAILED,
	}
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
