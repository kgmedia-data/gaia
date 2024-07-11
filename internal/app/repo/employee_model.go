package repo

import (
	"time"

	"github.com/kgmedia-data/gaia/internal/app/domain"
)

type Employee struct {
	Base
	Id        int       `gorm:"column:id"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	BirthDate time.Time `gorm:"column:birth_date"`
}

func (e Employee) TableName() string {
	return "employee"
}

func NewEmployee(employee domain.Employee) Employee {
	return Employee{
		Id:        employee.Id,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		BirthDate: employee.BirthDate,
	}
}

func toDomainEmployee(e Employee) domain.Employee {
	return domain.Employee{
		Id:        e.Id,
		FirstName: e.FirstName,
		LastName:  e.LastName,
		BirthDate: e.BirthDate,
	}
}

func toDomainEmployees(employees []Employee) []domain.Employee {
	return convertSlice(employees, toDomainEmployee)
}
