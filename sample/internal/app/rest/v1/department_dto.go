package v1

import "github.com/kgmedia-data/gaia/sample/internal/app/domain"

type departmentReqDto struct {
	Name string `json:"name"`
}

func (d departmentReqDto) toDomain() domain.Department {
	return domain.Department{
		Name: d.Name,
	}
}

type departmentRespDto struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func NewDepartmentRespDto(d domain.Department) departmentRespDto {
	return departmentRespDto{
		Id:   d.Id,
		Name: d.Name,
	}
}

func NewDepartmentsRespDto(departments []domain.Department) []departmentRespDto {
	return convertSlice(departments, NewDepartmentRespDto)
}

func (d departmentRespDto) toDomain() domain.Department {
	return domain.Department{
		Id:   d.Id,
		Name: d.Name,
	}
}
