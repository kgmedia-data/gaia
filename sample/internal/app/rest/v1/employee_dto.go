package v1

import (
	"time"

	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
)

type employeeDto struct {
	Id             int    `json:"id,omitempty"`
	EmployeeNumber string `json:"employeeNumber"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	BirthDate      string `json:"birthDate"`
}

func NewEmployeeDto(e domain.Employee) employeeDto {
	return employeeDto{
		Id:             e.Id,
		EmployeeNumber: e.EmployeeNumber,
		FirstName:      e.FirstName,
		LastName:       e.LastName,
		BirthDate:      e.BirthDate.Format("2006-01-02"),
	}
}

func NewEmployeesDto(employees []domain.Employee) []employeeDto {
	return convertSlice(employees, NewEmployeeDto)
}

func (e employeeDto) toDomain() domain.Employee {
	birthDate, _ := time.Parse("2006-01-02", e.BirthDate)

	return domain.Employee{
		Id:             e.Id,
		EmployeeNumber: e.EmployeeNumber,
		FirstName:      e.FirstName,
		LastName:       e.LastName,
		BirthDate:      birthDate,
	}
}
