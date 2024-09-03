package repo

import (
	"fmt"
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

func (s *DepartmentSuite) TearDownTest() {
	fmt.Println("teardown")
	truncateTables(s.repo.GormDB, "department")
}

func (s *DepartmentSuite) TearDownSuite() {
	truncateTables(s.repo.GormDB, "department")
	dd, _ := s.repo.GormDB.DB()
	dd.Close()
}

func TestDepartmentRepoSuite(t *testing.T) {
	suite.Run(t, new(DepartmentSuite))
}

func (s *DepartmentSuite) TestDepartmentRepo_GetDepartment() {
	type (
		args struct {
			id int
		}
	)

	testCases := []struct {
		desc      string
		args      args
		wantErr   bool
		checkFunc func(args args, dept domain.Department)
	}{
		{
			desc:    "success get department",
			wantErr: false,
			args: args{
				id: 1,
			},
			checkFunc: func(args args, dept domain.Department) {
				s.Equal(args.id, dept.Id)
				s.Equal("IT", dept.Name)
			},
		}, {
			desc:    "failed get department not found",
			wantErr: true,
			args: args{
				id: 2,
			},
			checkFunc: func(args args, dept domain.Department) {
				s.Zero(dept)
			},
		},
	}

	err := s.repo.GormDB.Create(&Department{Name: "IT"}).Error
	s.NoError(err)

	for _, tc := range testCases {
		s.Run(tc.desc, func() {
			dept, err := s.repo.GetDepartment(tc.args.id)
			if (err != nil) != tc.wantErr {
				s.Fail("error is not as expected")
			}
			tc.checkFunc(tc.args, dept)
		})
	}
}

func (s *DepartmentSuite) TestDepartmentRepo_GetDepartments() {
	type (
		args struct {
			offset int
			limit  int
		}
	)

	testCases := []struct {
		desc      string
		args      args
		wantErr   bool
		checkFunc func(args args, depts []domain.Department)
	}{
		{
			desc:    "success get departments",
			wantErr: false,
			args: args{
				offset: 0,
				limit:  100,
			},
			checkFunc: func(args args, depts []domain.Department) {
				s.Len(depts, 3)
				s.Equal("IT", depts[0].Name)
				s.Equal("HR", depts[1].Name)
				s.Equal("Finance", depts[2].Name)
			},
		}, {
			desc:    "success get departments with offset and limit",
			wantErr: false,
			args: args{
				offset: 1,
				limit:  2,
			},
			checkFunc: func(args args, depts []domain.Department) {
				s.Len(depts, 2)
				s.Equal("HR", depts[0].Name)
				s.Equal("Finance", depts[1].Name)
			},
		},
	}

	err := s.repo.GormDB.Create(&Department{Name: "IT"}).Error
	s.NoError(err)
	err = s.repo.GormDB.Create(&Department{Name: "HR"}).Error
	s.NoError(err)
	err = s.repo.GormDB.Create(&Department{Name: "Finance"}).Error
	s.NoError(err)

	for _, tc := range testCases {
		s.Run(tc.desc, func() {
			depts, err := s.repo.GetDepartments(tc.args.offset, tc.args.limit)
			if (err != nil) != tc.wantErr {
				s.Fail("error is not as expected")
			}
			tc.checkFunc(tc.args, depts)
		})
	}
}

func (s *DepartmentSuite) TestDepartmentRepo_InsertDepartment() {
	type (
		args struct {
			dept domain.Department
		}
	)

	testCases := []struct {
		desc      string
		args      args
		wantErr   bool
		checkFunc func(args args)
	}{
		{
			desc:    "success insert department",
			wantErr: false,
			args: args{
				dept: domain.Department{
					Name: "IT",
				},
			},
			checkFunc: func(args args) {
				s.NotZero(args.dept.Id)
				var dept Department
				err := s.repo.GormDB.Where("id = ?", args.dept.Id).First(&dept).Error
				s.NoError(err)
				s.Equal(args.dept.Name, dept.Name)
				s.Equal(false, dept.IsDeleted)
				s.Equal(dept.CreatedAt, dept.UpdatedAt)
			},
		}, {
			desc:    "failed insert duplicate department",
			wantErr: true,
			args: args{
				dept: domain.Department{
					Name: "IT",
				},
			},
			checkFunc: func(args args) {},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.desc, func() {
			dept, err := s.repo.InsertDepartment(tc.args.dept)
			if (err != nil) != tc.wantErr {
				s.Fail("error is not as expected")
			}
			tc.checkFunc(args{dept})
		})
	}
}

func (s *DepartmentSuite) TestDepartmentRepo_UpdateDepartment() {
	type (
		args struct {
			dept domain.Department
		}
	)

	testCases := []struct {
		desc      string
		args      args
		wantErr   bool
		checkFunc func(args args)
	}{
		{
			desc:    "success update department",
			wantErr: false,
			args: args{
				dept: domain.Department{
					Id:   1,
					Name: "Information Technology",
				},
			},
			checkFunc: func(args args) {
				s.NotZero(args.dept.Id)
				var dept Department
				err := s.repo.GormDB.Where("id = ?", args.dept.Id).First(&dept).Error
				s.NoError(err)
				s.Equal(args.dept.Name, dept.Name)
				s.Equal(false, dept.IsDeleted)
				s.NotEqual(dept.CreatedAt, dept.UpdatedAt)
			},
		}, {
			desc:    "failed insert duplicate department",
			wantErr: true,
			args: args{
				dept: domain.Department{
					Id:   1,
					Name: "HR",
				},
			},
			checkFunc: func(args args) {},
		},
	}

	err := s.repo.GormDB.Create(&Department{Name: "IT"}).Error
	s.NoError(err)
	err = s.repo.GormDB.Create(&Department{Name: "HR"}).Error
	s.NoError(err)

	for _, tc := range testCases {
		s.Run(tc.desc, func() {
			dept, err := s.repo.UpdateDepartment(tc.args.dept)
			if (err != nil) != tc.wantErr {
				s.Fail("error is not as expected")
			}
			tc.checkFunc(args{dept})
		})
	}
}

func (s *DepartmentSuite) TestDepartmentRepo_DeleteDepartment() {
	type (
		args struct {
			id int
		}
	)

	testCases := []struct {
		desc      string
		args      args
		wantErr   bool
		checkFunc func(args args)
	}{
		{
			desc:    "success delete department",
			wantErr: false,
			args: args{
				id: 1,
			},
			checkFunc: func(args args) {
				var dept Department
				err := s.repo.GormDB.Where("id = ?", args.id).First(&dept).Error
				s.NoError(err)
				s.True(dept.IsDeleted)
			},
		}, {
			desc:    "success delete non existent department",
			wantErr: false,
			args: args{
				id: 2,
			},
			checkFunc: func(args args) {},
		},
	}

	err := s.repo.GormDB.Create(&Department{Name: "IT"}).Error
	s.NoError(err)

	for _, tc := range testCases {
		s.Run(tc.desc, func() {
			err := s.repo.DeleteDepartment(tc.args.id)
			if (err != nil) != tc.wantErr {
				s.Fail("error is not as expected")
			}
			tc.checkFunc(tc.args)
		})
	}
}
