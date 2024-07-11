package service

import (
	"fmt"

	"github.com/kgmedia-data/gaia/internal/app/repo"
)

type EmployeeService struct {
	employeeRepo repo.IEmployeeRepo
}

func NewEmployeeService(employeeRepo repo.IEmployeeRepo) *EmployeeService {
	return &EmployeeService{
		employeeRepo: employeeRepo,
	}
}

func (s EmployeeService) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("EmployeeService.(%v)(%v) %w", method, params, err)
}
