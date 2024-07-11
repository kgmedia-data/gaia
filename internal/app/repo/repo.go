package repo

import "github.com/kgmedia-data/gaia/internal/app/domain"

type IEmployeeRepo interface {
	GetEmployee(id int) (domain.Employee, error)
	GetEmployees() ([]domain.Employee, error)
	InsertEmployee(employee domain.Employee) (domain.Employee, error)
	UpdateEmployee(employee domain.Employee) (domain.Employee, error)
	DeleteEmployee(id int) error
}
