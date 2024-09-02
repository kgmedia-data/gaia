package service

import (
	"fmt"

	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
	"github.com/kgmedia-data/gaia/sample/internal/app/repo"
)

type EmployeeService struct {
	deptRepo     repo.IDepartmentRepo
	employeeRepo repo.IEmployeeRepo
}

func NewEmployeeService(employeeRepo repo.IEmployeeRepo,
	deptRepo repo.IDepartmentRepo) *EmployeeService {
	return &EmployeeService{
		employeeRepo: employeeRepo,
		deptRepo:     deptRepo,
	}
}

func (s EmployeeService) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("EmployeeService.(%v)(%v) %w", method, params, err)
}

func (s EmployeeService) GetEmployee(deptId, id int) (domain.Employee, error) {
	employee, err := s.employeeRepo.GetEmployee(id)
	if err != nil {
		return domain.Employee{}, s.error(err, "GetEmployee", id)
	}

	if employee.Department.Id != deptId {
		return domain.Employee{}, s.error(fmt.Errorf("employee not found"),
			"GetEmployee", deptId, id)
	}

	return employee, nil
}

func (s EmployeeService) GetEmployees(deptId, offset, limit int) ([]domain.Employee, error) {
	employees, err := s.employeeRepo.GetEmployees(offset, limit, deptId)
	if err != nil {
		return []domain.Employee{}, s.error(err, "GetEmployees", deptId, offset, limit)
	}
	return employees, nil
}

func (s EmployeeService) InsertEmployee(deptId int,
	employee domain.Employee) (domain.Employee, error) {

	dept, err := s.deptRepo.GetDepartment(deptId)
	if err != nil {
		return domain.Employee{}, s.error(err, "InsertEmployee", deptId, employee)
	}

	employee.Department = dept
	if !employee.IsValid() {
		return domain.Employee{}, s.error(fmt.Errorf("invalid employee"),
			"InsertEmployee", deptId, employee)
	}

	employee, err = s.employeeRepo.InsertEmployee(employee)
	if err != nil {
		return domain.Employee{}, s.error(err, "InsertEmployee", deptId, employee)
	}

	return employee, nil
}

func (s EmployeeService) UpdateEmployee(deptId int, employee domain.Employee) (domain.Employee, error) {
	dept, err := s.deptRepo.GetDepartment(deptId)
	if err != nil {
		return domain.Employee{}, s.error(err, "UpdateEmployee", deptId, employee)
	}

	employee.Department = dept
	if !employee.IsValid() || employee.Id == 0 {
		return domain.Employee{}, s.error(fmt.Errorf("invalid employee"),
			"UpdateEmployee", deptId, employee)
	}

	_, err = s.GetEmployee(deptId, employee.Id)
	if err != nil {
		return domain.Employee{}, s.error(err, "UpdateEmployee", deptId, employee)
	}

	employee, err = s.employeeRepo.UpdateEmployee(employee)
	if err != nil {
		return domain.Employee{}, s.error(err, "UpdateEmployee", deptId, employee)
	}

	return employee, nil
}

func (s EmployeeService) DeleteEmployee(deptId, id int) error {
	if id == 0 {
		return s.error(fmt.Errorf("invalid employee id"), "DeleteEmployee", deptId, id)
	}

	_, err := s.GetEmployee(deptId, id)
	if err != nil {
		return s.error(err, "DeleteEmployee", deptId, id)
	}

	err = s.employeeRepo.DeleteEmployee(id)
	if err != nil {
		return s.error(err, "DeleteEmployee", deptId, id)
	}

	return nil
}
