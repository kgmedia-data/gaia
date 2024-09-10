package repo

import "github.com/kgmedia-data/gaia/sample/internal/app/domain"

type DepartmentMemRepo struct {
	Departments map[int]domain.Department
}

func NewDepartmentMemRepo() *DepartmentMemRepo {
	deptMap := make(map[int]domain.Department)
	deptMap[1] = domain.Department{
		Id:   1,
		Name: "IT",
	}

	return &DepartmentMemRepo{
		Departments: deptMap,
	}
}

func (r *DepartmentMemRepo) GetDepartment(id int) (domain.Department, error) {
	dept, ok := r.Departments[id]
	if !ok {
		return domain.Department{}, nil
	}
	return dept, nil
}

func (r *DepartmentMemRepo) GetDepartments(offset, limit int) ([]domain.Department, error) {
	var depts []domain.Department
	for _, dept := range r.Departments {
		depts = append(depts, dept)
	}
	return depts, nil
}

func (r *DepartmentMemRepo) InsertDepartment(department domain.Department) (domain.Department, error) {
	department.Id = len(r.Departments) + 1
	r.Departments[department.Id] = department
	return department, nil
}

func (r *DepartmentMemRepo) UpdateDepartment(department domain.Department) (domain.Department, error) {
	r.Departments[department.Id] = department
	return department, nil
}

func (r *DepartmentMemRepo) DeleteDepartment(id int) error {
	delete(r.Departments, id)
	return nil
}
