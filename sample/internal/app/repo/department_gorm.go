package repo

import (
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
)

type DepartmentGormRepo struct {
	GormRepo
}

func NewDepartmentGorm(gormRepo *GormRepo) *DepartmentGormRepo {
	return &DepartmentGormRepo{
		GormRepo: *gormRepo,
	}
}

func (r DepartmentGormRepo) error(err error, method string, params ...interface{}) error {
	message := fmt.Errorf("DepartmentGormRepo.(%v)(%v) %w", method, params, err)
	logrus.Error(message)
	return message
}

func (r *DepartmentGormRepo) GetDepartment(id int) (domain.Department, error) {
	department := Department{}
	if err := r.GormDB.Where("is_deleted = ?", false).First(&department, id).Error; err != nil {
		return domain.Department{}, r.error(err, "GetDepartment", id)
	}
	return department.toDomain(), nil
}

func (r *DepartmentGormRepo) GetDepartments(offset, limit int) ([]domain.Department, error) {
	var departments Departments
	logrus.WithFields(logrus.Fields{
		"gcp":       true,
		"firstName": "Kompas",
		"lastName":  "Gramedia",
	}).Info("GetDepartments Repo is run")
	tx := r.GormDB.
		Where("is_deleted = ?", false).
		Offset(offset).
		Limit(limit).
		Order("id ASC")

	if err := tx.Find(&departments).Error; err != nil {
		return []domain.Department{}, r.error(err, "GetDepartments")
	}
	return departments.toDomain(), nil
}

func (r *DepartmentGormRepo) InsertDepartment(d domain.Department) (domain.Department, error) {
	department := NewDepartment(d)
	if err := r.GormDB.Create(&department).Error; err != nil {
		return domain.Department{}, r.error(err, "InsertDepartment", department)
	}
	return department.toDomain(), nil
}

func (r *DepartmentGormRepo) UpdateDepartment(d domain.Department) (domain.Department, error) {
	department := NewDepartment(d)
	if err := r.GormDB.Omit("created_at").Save(&department).Error; err != nil {
		return domain.Department{}, r.error(err, "UpdateDepartment", department)
	}
	return department.toDomain(), nil
}

func (r *DepartmentGormRepo) DeleteDepartment(id int) error {
	err := r.GormDB.Model(&Department{}).Where("id = ?", id).Update("is_deleted", true).Error
	if err != nil {
		return r.error(err, "DeleteDepartment", id)
	}
	return nil
}
