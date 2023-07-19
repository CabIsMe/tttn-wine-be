package services

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
)

type CustomerService interface {
	GetCustomerInfoService(customerId string) (*models.Customer, *internal.SystemStatus)
	UpdateCustomerService(customer models.Customer) *internal.SystemStatus
}
type customer_service struct {
	rp repositories.Repos
}

func NewCustomerService(rp repositories.Repos) CustomerService {
	return &customer_service{
		rp,
	}
}
func (s *customer_service) GetCustomerInfoService(customerId string) (*models.Customer, *internal.SystemStatus) {
	result, err := s.rp.GetCustomer(customerId)
	if err != nil {
		return nil, internal.SysStatus.DbFailed
	}
	if result == nil {
		return nil, &internal.SystemStatus{
			Status: internal.CODE_DB_FAILED,
			Msg:    internal.MSG_DB_FAILED,
			Detail: "Customer not found",
		}
	}
	return result, nil
}
func (s *customer_service) UpdateCustomerService(customer models.Customer) *internal.SystemStatus {
	err := s.rp.UpdateCustomer(customer)
	if err != nil {
		return internal.SysStatus.DbFailed
	}
	return nil
}
