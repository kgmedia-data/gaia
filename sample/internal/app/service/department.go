package service

import (
	"fmt"

	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
	"github.com/kgmedia-data/gaia/sample/internal/app/repo"
)

type DepartmentService struct {
	deptRepo repo.IDepartmentRepo
}

func NewDepartmentService(deptRepo repo.IDepartmentRepo) *DepartmentService {
	return &DepartmentService{deptRepo: deptRepo}
}

func (s DepartmentService) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("DepartmentService.(%v)(%v) %w", method, params, err)
}

func (s DepartmentService) GetDepartment(id int) (domain.Department, error) {
	dept, err := s.deptRepo.GetDepartment(id)
	if err != nil {
		return domain.Department{}, s.error(err, "GetDepartment", id)
	}
	return dept, nil
}

func (s DepartmentService) GetDepartments(offset, limit int) ([]domain.Department, error) {
	depts, err := s.deptRepo.GetDepartments(offset, limit)
	if err != nil {
		return []domain.Department{}, s.error(err, "GetDepartments", offset, limit)
	}
	return depts, nil
}

func (s DepartmentService) InsertDepartment(dept domain.Department) (domain.Department, error) {
	if !dept.IsValid() {
		return domain.Department{}, s.error(fmt.Errorf("invalid department"),
			"InsertDepartment", dept)
	}

	dept, err := s.deptRepo.InsertDepartment(dept)
	if err != nil {
		return domain.Department{}, s.error(err, "InsertDepartment", dept)
	}
	return dept, nil
}

func (s DepartmentService) UpdateDepartment(dept domain.Department) (domain.Department, error) {
	if !dept.IsValid() || dept.Id == 0 {
		return domain.Department{}, s.error(fmt.Errorf("invalid department"),
			"UpdateDepartment", dept)
	}

	_, err := s.GetDepartment(dept.Id)
	if err != nil {
		return domain.Department{}, s.error(err, "UpdateDepartment", dept)
	}

	dept, err = s.deptRepo.UpdateDepartment(dept)
	if err != nil {
		return domain.Department{}, s.error(err, "UpdateDepartment", dept)
	}
	return dept, nil
}

func (s DepartmentService) DeleteDepartment(id int) error {
	if id == 0 {
		return s.error(fmt.Errorf("invalid department id"), "DeleteDepartment", id)
	}

	_, err := s.GetDepartment(id)
	if err != nil {
		return s.error(err, "DeleteDepartment", id)
	}

	err = s.deptRepo.DeleteDepartment(id)
	if err != nil {
		return s.error(err, "DeleteDepartment", id)
	}
	return nil
}
