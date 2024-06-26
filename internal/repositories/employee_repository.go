package repositories

import (
	"errors"
	"fmt"

	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	GetEmployee(empId string) (*models.Employee, error)
	GetEmployeeByEmail(email string) (*models.Employee, error)
	GetAccountInfo(username string) (*models.AccountInfo, error)
	GetAllDeliverers() ([]*models.Deliverer, error)
}

type employee_repos struct {
}

func NewEmployeeRepository() EmployeeRepository {
	return &employee_repos{}
}

func (r *employee_repos) GetEmployee(empId string) (*models.Employee, error) {
	model := &models.Employee{}
	result := internal.Db.Where(fmt.Sprintf("%s = ?", model.ColumnEmployeeId()), empId).
		First(&model).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		internal.Log.Error("GetEmployee", zap.Any("Error query", result))
		return nil, nil
	}
	return model, result
}
func (r *employee_repos) GetEmployeeByEmail(email string) (*models.Employee, error) {
	model := &models.Employee{}
	result := internal.Db.Where(fmt.Sprintf("%s = ?", model.ColumnEmail()), email).First(&model).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		internal.Log.Error("GetEmployeeByEmail", zap.Any("Error query", result))
		return nil, nil
	}
	return model, result
}

func (r *employee_repos) GetAccountInfo(username string) (*models.AccountInfo, error) {
	model := &models.AccountInfo{}
	result := internal.Db.Table("accounts").
		Preload("RoleInfo").
		Preload("EmployeeInfo").
		Where(fmt.Sprintf("%s = ?", "accounts.username"), username).
		First(&model).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		internal.Log.Error("GetAccountInfo", zap.Any("Error query", result))
		return nil, nil
	}
	return model, result
}
func (r *employee_repos) GetAllDeliverers() ([]*models.Deliverer, error) {
	model := &models.Employee{}
	account := &models.Account{}
	var employees []*models.Deliverer
	err := internal.Db.Debug().Table(model.TableName()).
		Select("*").
		Joins(fmt.Sprintf("INNER JOIN %s ON %s.%s = %s.%s", account.TableName(),
			account.TableName(), account.ColumnUsername(), model.TableName(), model.ColumnEmail())). //INNER JOIN accounts ON accounts.username = employee.email
		Where(fmt.Sprintf("%s.%s = ?", account.TableName(), account.ColumnRoleId()), 3).
		Find(&employees).
		Error
	if err != nil {
		internal.Log.Error("GetAllDeliverers", zap.Any("Error Get ListDeliverers", err))
		return nil, err
	}
	for _, record := range employees {
		err = internal.Db.Debug().
			Model(&models.CustomerOrder{}).
			Where("deliverer_id = ? AND status = ?", record.EmployeeId, models.Cos.ORDER_CONFIRM.StatusCode).
			Count(&record.NumberOfDelivered).Error
		if err != nil {
			internal.Log.Error("GetAllDeliverers", zap.Any("Error Get NumberOfDelivered", err))
			return nil, err
		}
	}
	return employees, err
}
