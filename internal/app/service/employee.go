package service

import (
	"fmt"

	"github.com/kgmedia-data/gaia/internal/app/domain"
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

func (s *EmployeeService) GetEmployee(id int) (domain.Employee, error) {
	employee, err := s.employeeRepo.GetEmployee(id)
	if err != nil {
		return domain.Employee{}, s.error(err, "GetEmployee", id)
	}
	return employee, nil
}

func (s *EmployeeService) GetEmployees() ([]domain.Employee, error) {
	employees, err := s.employeeRepo.GetEmployees()
	if err != nil {
		return []domain.Employee{}, s.error(err, "GetEmployees")
	}
	return employees, nil
}

func (s *EmployeeService) InsertEmployee(employee domain.Employee) (domain.Employee, error) {
	if !employee.IsValid() {
		return domain.Employee{}, s.error(fmt.Errorf("invalid employee"),
			"InsertEmployee", employee)
	}

	employee, err := s.employeeRepo.InsertEmployee(employee)
	if err != nil {
		return domain.Employee{}, s.error(err, "InsertEmployee", employee)
	}
	return employee, nil
}

func (s *EmployeeService) UpdateEmployee(employee domain.Employee) (domain.Employee, error) {
	if !employee.IsValid() {
		return domain.Employee{}, s.error(fmt.Errorf("invalid employee"),
			"UpdateEmployee", employee)
	}

	_, err := s.employeeRepo.GetEmployee(employee.Id)
	if err != nil || employee.Id == 0 {
		return domain.Employee{}, s.error(fmt.Errorf("employee does not exists"), "UpdateEmployee", employee)
	}

	employee, err = s.employeeRepo.UpdateEmployee(employee)
	if err != nil {
		return domain.Employee{}, s.error(err, "UpdateEmployee", employee)
	}
	return employee, nil
}

func (s *EmployeeService) DeleteEmployee(id int) error {
	_, err := s.employeeRepo.GetEmployee(id)
	if err != nil || id == 0 {
		return s.error(fmt.Errorf("employee does not exists"),
			"DeleteEmployee", id)
	}

	err = s.employeeRepo.DeleteEmployee(id)
	if err != nil {
		return s.error(err, "DeleteEmployee", id)
	}
	return nil
}
