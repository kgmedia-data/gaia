package repo

import (
	"log"
	"testing"

	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
	"github.com/stretchr/testify/suite"
)

type DepartmentSuite struct {
	suite.Suite
	repo *DepartmentGormRepo
}

func (s *DepartmentSuite) SetupSuite() {
	db, err := NewGormRepo(TEST_CONN_STRING, 1)
	if err != nil {
		log.Fatalln(err)
	}
	truncateTables(db.GormDB, "department")

	s.repo = NewDepartmentGorm(db)
}

func (s *DepartmentSuite) TearDownSuite() {
	truncateTables(s.repo.GormDB, "department")
	dd, _ := s.repo.GormDB.DB()
	dd.Close()
}

func TestDepartmentRepoSuite(t *testing.T) {
	suite.Run(t, new(DepartmentSuite))
}

func (s *DepartmentSuite) TestDepartmentRepo() {
	depts, err := s.repo.GetDepartments(0, 100)
	s.NoError(err)
	s.Len(depts, 0)

	deptIt := domain.Department{
		Name: "IT",
	}
	deptIt, err = s.repo.InsertDepartment(deptIt)
	s.NoError(err)
	s.NotZero(deptIt.Id)

	deptHr := domain.Department{
		Name: "HR",
	}
	deptHr, err = s.repo.InsertDepartment(deptHr)
	s.NoError(err)
	s.NotZero(deptHr.Id)

	depts, err = s.repo.GetDepartments(0, 100)
	s.NoError(err)
	s.Len(depts, 2)
	s.Equal(deptIt.Name, depts[0].Name)
	s.Equal(deptHr.Name, depts[1].Name)

	result, err := s.repo.GetDepartment(deptIt.Id)
	s.NoError(err)
	s.Equal(deptIt.Name, result.Name)

	deptIt.Name = "Information Technology"
	deptIt, err = s.repo.UpdateDepartment(deptIt)
	s.NoError(err)

	result, err = s.repo.GetDepartment(deptIt.Id)
	s.NoError(err)
	s.Equal(deptIt.Name, result.Name)

	err = s.repo.DeleteDepartment(deptIt.Id)
	s.NoError(err)

	depts, err = s.repo.GetDepartments(0, 100)
	s.NoError(err)
	s.Len(depts, 1)
	s.Equal(deptHr.Name, depts[0].Name)

	result, err = s.repo.GetDepartment(deptIt.Id)
	s.Error(err)
}
