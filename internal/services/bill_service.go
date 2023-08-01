package services

import (
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type BillService interface {
	GetBillByCustomerOrderIdService(customerOrderId string) (*models.Bill, *internal.SystemStatus)
	CreateBillService(bill models.Bill) *internal.SystemStatus
}
type bill_service struct {
	rp repositories.Repos
}

func NewBillService(rp repositories.Repos) BillService {
	return &bill_service{
		rp,
	}
}

func (s *bill_service) GetBillByCustomerOrderIdService(customerOrderId string) (*models.Bill, *internal.SystemStatus) {
	model, err := s.rp.GetBillByCustomerOrderId(customerOrderId)
	if err != nil {
		return nil, internal.SysStatus.DbFailed
	}
	return model, nil
}

func (s *bill_service) CreateBillService(bill models.Bill) *internal.SystemStatus {
	errResult := internal.SystemStatus{
		Status: internal.CODE_DB_FAILED,
		Msg:    internal.MSG_DB_FAILED,
	}
	nanoId, _ := gonanoid.New()
	bill.BillId = nanoId
	err := s.rp.CreateBill(bill)
	if err != nil {
		errResult.Detail = err.Error()
		return &errResult
	}
	return nil
}
