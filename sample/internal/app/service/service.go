package service

import "github.com/kgmedia-data/gaia/sample/internal/app/domain"

type IDepartmentService interface {
	GetDepartment(id int) (domain.Department, error)
	GetDepartments(offset, limit int) ([]domain.Department, error)
	InsertDepartment(department domain.Department) (domain.Department, error)
	UpdateDepartment(department domain.Department) (domain.Department, error)
	DeleteDepartment(id int) error
}

type IEmployeeService interface {
	GetEmployee(deptId, id int) (domain.Employee, error)
	GetEmployees(deptId, offset, limit int) ([]domain.Employee, error)
	InsertEmployee(deptId int, employee domain.Employee) (domain.Employee, error)
	UpdateEmployee(deptId int, employee domain.Employee) (domain.Employee, error)
	DeleteEmployee(deptId, id int) error
}
