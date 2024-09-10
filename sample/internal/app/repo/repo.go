package repo

import "github.com/kgmedia-data/gaia/sample/internal/app/domain"

type IEmployeeRepo interface {
	GetEmployee(id int) (domain.Employee, error)
	GetEmployees(offset, limit int, departmentId ...int) ([]domain.Employee, error)
	InsertEmployee(employee domain.Employee) (domain.Employee, error)
	UpdateEmployee(employee domain.Employee) (domain.Employee, error)
	DeleteEmployee(id int) error
}

type IDepartmentRepo interface {
	GetDepartment(id int) (domain.Department, error)
	GetDepartments(offset, limit int) ([]domain.Department, error)
	InsertDepartment(department domain.Department) (domain.Department, error)
	UpdateDepartment(department domain.Department) (domain.Department, error)
	DeleteDepartment(id int) error
}
