package repo

import (
	"time"

	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
)

type Employee struct {
	Base
	Id             int       `gorm:"column:id"`
	EmployeeNumber string    `gorm:"column:employee_number"`
	FirstName      string    `gorm:"column:first_name"`
	LastName       string    `gorm:"column:last_name"`
	BirthDate      time.Time `gorm:"column:birth_date"`
	DepartmentId   int       `gorm:"column:department_id"`
	Department     Department
}

type Employees []Employee

func (e Employee) TableName() string {
	return "employee"
}

func NewEmployee(employee domain.Employee) Employee {
	return Employee{
		Id:             employee.Id,
		EmployeeNumber: employee.EmployeeNumber,
		FirstName:      employee.FirstName,
		LastName:       employee.LastName,
		BirthDate:      employee.BirthDate,
		DepartmentId:   employee.Department.Id,
		Department:     NewDepartment(employee.Department),
	}
}

func (e Employee) toDomain() domain.Employee {
	return domain.Employee{
		Id:             e.Id,
		EmployeeNumber: e.EmployeeNumber,
		FirstName:      e.FirstName,
		LastName:       e.LastName,
		BirthDate:      e.BirthDate,
		Department:     e.Department.toDomain(),
	}
}

func (es Employees) toDomain() []domain.Employee {
	var emps []domain.Employee
	for _, e := range es {
		emps = append(emps, e.toDomain())
	}
	return emps
}
