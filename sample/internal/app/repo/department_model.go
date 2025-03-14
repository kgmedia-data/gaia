package repo

import (
	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
)

type Department struct {
	Base
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

type Departments []Department

func (e Department) TableName() string {
	return "department"
}

func NewDepartment(d domain.Department) Department {
	return Department{
		Id:   d.Id,
		Name: d.Name,
	}
}

func (d Department) toDomain() domain.Department {
	return domain.Department{
		Id:   d.Id,
		Name: d.Name,
	}
}

func (ds Departments) toDomain() []domain.Department {
	var depts []domain.Department
	for _, d := range ds {
		depts = append(depts, d.toDomain())
	}
	return depts
}
