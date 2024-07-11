package service

import "github.com/kgmedia-data/gaia/internal/app/domain"

type IEmployeeService interface {
	GetEmployee(id int) (domain.Employee, error)
	GetEmployees() ([]domain.Employee, error)
	InsertEmployee(employee domain.Employee) (domain.Employee, error)
	UpdateEmployee(employee domain.Employee) (domain.Employee, error)
	DeleteEmployee(id int) error
}
