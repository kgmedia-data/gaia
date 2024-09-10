package repo

import "github.com/kgmedia-data/gaia/sample/internal/app/domain"

type EmployeeMemRepo struct {
	Employees map[int]domain.Employee
}

func NewEmployeeMemRepo() *EmployeeMemRepo {
	return &EmployeeMemRepo{
		Employees: make(map[int]domain.Employee),
	}
}

func (r *EmployeeMemRepo) GetEmployee(id int) (domain.Employee, error) {
	employee, ok := r.Employees[id]
	if !ok {
		return domain.Employee{}, nil
	}
	return employee, nil
}

func (r *EmployeeMemRepo) GetEmployees(offset, limit int, departmentId ...int) (
	[]domain.Employee, error) {
	var employees []domain.Employee
	for _, employee := range r.Employees {
		employees = append(employees, employee)
	}
	return employees, nil
}

func (r *EmployeeMemRepo) InsertEmployee(employee domain.Employee) (domain.Employee, error) {
	employee.Id = len(r.Employees) + 1
	r.Employees[employee.Id] = employee
	return employee, nil
}

func (r *EmployeeMemRepo) UpdateEmployee(employee domain.Employee) (domain.Employee, error) {
	r.Employees[employee.Id] = employee
	return employee, nil
}

func (r *EmployeeMemRepo) DeleteEmployee(id int) error {
	delete(r.Employees, id)
	return nil
}
