package v1

import (
	"time"

	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
)

type employeeReqDto struct {
	EmployeeNumber string `json:"employeeNumber"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	BirthDate      string `json:"birthDate"`
}

func (e employeeReqDto) toDomain() domain.Employee {
	birthDate, _ := time.Parse("2006-01-02", e.BirthDate)

	return domain.Employee{
		EmployeeNumber: e.EmployeeNumber,
		FirstName:      e.FirstName,
		LastName:       e.LastName,
		BirthDate:      birthDate,
	}
}

type employeeRespDto struct {
	Id             int    `json:"id,omitempty"`
	EmployeeNumber string `json:"employeeNumber"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	BirthDate      string `json:"birthDate" example:"2006-01-02"`
}

func NewEmployeeRespDto(e domain.Employee) employeeRespDto {
	return employeeRespDto{
		Id:             e.Id,
		EmployeeNumber: e.EmployeeNumber,
		FirstName:      e.FirstName,
		LastName:       e.LastName,
		BirthDate:      e.BirthDate.Format("2006-01-02"),
	}
}

func NewEmployeesRespDto(employees []domain.Employee) []employeeRespDto {
	return convertSlice(employees, NewEmployeeRespDto)
}

func (e employeeRespDto) toDomain() domain.Employee {
	birthDate, _ := time.Parse("2006-01-02", e.BirthDate)

	return domain.Employee{
		Id:             e.Id,
		EmployeeNumber: e.EmployeeNumber,
		FirstName:      e.FirstName,
		LastName:       e.LastName,
		BirthDate:      birthDate,
	}
}
