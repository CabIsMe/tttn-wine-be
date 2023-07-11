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
}

type employee_repos struct {
}

func NewEmployeeRepository() EmployeeRepository {
	return &employee_repos{}
}

func (r *employee_repos) GetEmployee(empId string) (*models.Employee, error) {
	model := &models.Employee{}
	result := internal.Db.Where(fmt.Sprintf("%s = ?", model.ColumnEmployeeId()), empId).First(&model).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		internal.Log.Error("GetAccount", zap.Any("Error query", result))
		return nil, nil
	}
	return model, result
}
