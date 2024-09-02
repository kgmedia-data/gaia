package v1

import "github.com/kgmedia-data/gaia/sample/internal/app/domain"

type departmentDto struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func NewDepartmentDto(d domain.Department) departmentDto {
	return departmentDto{
		Id:   d.Id,
		Name: d.Name,
	}
}

func NewDepartmentsDto(departments []domain.Department) []departmentDto {
	return convertSlice(departments, NewDepartmentDto)
}

func toDomainDepartment(d departmentDto) domain.Department {
	return domain.Department{
		Id:   d.Id,
		Name: d.Name,
	}
}
