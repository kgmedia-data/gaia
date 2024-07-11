package repo

import (
	"fmt"

	"github.com/kgmedia-data/gaia/internal/app/domain"
)

type EmployeeGormRepo struct {
	GormRepo
}

func NewEmployeeGorm(gormRepo *GormRepo) *EmployeeGormRepo {
	return &EmployeeGormRepo{
		GormRepo: *gormRepo,
	}
}

func (r EmployeeGormRepo) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("EmployeeGormRepo.(%v)(%v) %w", method, params, err)
}

func (r *EmployeeGormRepo) GetEmployee(id int) (domain.Employee, error) {
	employee := Employee{}
	if err := r.GormDB.First(&employee, id).Error; err != nil {
		return domain.Employee{}, r.error(err, "GetEmployee", id)
	}
	return toDomainEmployee(employee), nil
}

func (r *EmployeeGormRepo) GetEmployees() ([]domain.Employee, error) {
	employees := []Employee{}
	if err := r.GormDB.Find(&employees).Error; err != nil {
		return []domain.Employee{}, r.error(err, "GetEmployees")
	}
	return toDomainEmployees(employees), nil
}

func (r *EmployeeGormRepo) InsertEmployee(e domain.Employee) (domain.Employee, error) {
	employee := NewEmployee(e)
	if err := r.GormDB.Create(&employee).Error; err != nil {
		return domain.Employee{}, r.error(err, "InsertEmployee", employee)
	}
	return toDomainEmployee(employee), nil
}

func (r *EmployeeGormRepo) UpdateEmployee(e domain.Employee) (domain.Employee, error) {
	employee := NewEmployee(e)
	if err := r.GormDB.Save(&employee).Error; err != nil {
		return domain.Employee{}, r.error(err, "UpdateEmployee", employee)
	}
	return toDomainEmployee(employee), nil
}

func (r *EmployeeGormRepo) DeleteEmployee(id int) error {
	err := r.GormDB.Model(&Employee{}).Where("id = ?", id).Update("is_deleted", true).Error
	if err != nil {
		return r.error(err, "DeleteEmployee", id)
	}
	return nil
}
