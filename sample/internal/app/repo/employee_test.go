package repo

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
	"github.com/stretchr/testify/suite"
)

type EmployeeSuite struct {
	suite.Suite
	repo      *EmployeeGormRepo
	employess Employees
}

func (s *EmployeeSuite) SetupSuite() {
	db, err := NewGormRepo(TEST_CONN_STRING, 2)
	if err != nil {
		log.Fatalln(err)
	}
	truncateTables(db.GormDB, "department", "employee")

	db.GormDB.Exec("INSERT INTO department (id, name) VALUES (1, 'IT')")
	db.GormDB.Exec("INSERT INTO department (id, name) VALUES (2, 'HR')")

	s.employess = Employees{
		{
			EmployeeNumber: "111",
			FirstName:      "John",
			LastName:       "Doe",
			BirthDate:      time.Date(1981, 1, 1, 0, 0, 0, 0, time.UTC),
			DepartmentId:   1,
			Department: Department{
				Id:   1,
				Name: "IT",
			},
		},
		{
			EmployeeNumber: "222",
			FirstName:      "Jane",
			LastName:       "Doe",
			BirthDate:      time.Date(1984, 1, 1, 0, 0, 0, 0, time.UTC),
			DepartmentId:   2,
			Department: Department{
				Id:   2,
				Name: "HR",
			},
		},
	}

	s.repo = NewEmployeeGorm(db)
}

func (s *EmployeeSuite) TearDownSuite() {
	truncateTables(s.repo.GormDB, "department", "employee")
	dd, _ := s.repo.GormDB.DB()
	dd.Close()
}

func TestEmployeeRepoSuite(t *testing.T) {
	suite.Run(t, new(EmployeeSuite))
}

func (s *EmployeeSuite) SetupTest() {
	fmt.Println("setup test")

	tx := s.repo.GormDB.Begin()
	err := tx.Create(&s.employess[0]).Error
	s.NoError(err)

	err = tx.Create(&s.employess[1]).Error
	s.NoError(err)

	tx.Exec("SELECT setval('employee_id_seq', 3, false)")

	tx.Commit()
}

func (s *EmployeeSuite) TearDownTest() {
	fmt.Println("teardown test")
	truncateTables(s.repo.GormDB, "employee")
}

func (s *EmployeeSuite) TestEmployeeRepo_GetEmployee() {
	type (
		args struct {
			id int
		}
	)

	testCases := []struct {
		desc      string
		args      args
		wantErr   bool
		checkFunc func(args args, employee domain.Employee)
	}{
		{
			desc:    "success get employee",
			wantErr: false,
			args: args{
				id: 1,
			},
			checkFunc: func(args args, employee domain.Employee) {
				s.Equal(args.id, employee.Id)
				s.Equal("111", employee.EmployeeNumber)
				s.Equal("John", employee.FirstName)
				s.Equal("Doe", employee.LastName)
				s.Equal(1981, employee.BirthDate.Year())
				s.Equal(1, employee.Department.Id)
				s.Equal("IT", employee.Department.Name)
			},
		}, {
			desc:    "failed get employee not found",
			wantErr: true,
			args: args{
				id: 3,
			},
			checkFunc: func(args args, employee domain.Employee) {},
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			employee, err := s.repo.GetEmployee(tC.args.id)
			if tC.wantErr {
				s.Error(err)
				return
			}

			s.NoError(err)
			tC.checkFunc(tC.args, employee)
		})
	}
}

func (s *EmployeeSuite) TestEmployeeRepo_GetEmployees() {
	type (
		args struct {
			offset  int
			limit   int
			deptsId []int
		}
	)

	testCases := []struct {
		desc      string
		args      args
		wantErr   bool
		checkFunc func(args args, employees []domain.Employee)
	}{
		{
			desc:    "success get employees",
			wantErr: false,
			args: args{
				offset: 0,
				limit:  100,
			},
			checkFunc: func(args args, employees []domain.Employee) {
				s.Len(employees, 2)
				s.Equal(s.employess[0].EmployeeNumber, employees[0].EmployeeNumber)
				s.Equal(s.employess[0].FirstName, employees[0].FirstName)
				s.Equal(s.employess[0].LastName, employees[0].LastName)
				s.Equal(s.employess[0].BirthDate.Year(), employees[0].BirthDate.Year())
				s.Equal(s.employess[0].Department.Id, employees[0].Department.Id)
				s.Equal(s.employess[0].Department.Name, employees[0].Department.Name)

				s.Equal(s.employess[1].EmployeeNumber, employees[1].EmployeeNumber)
				s.Equal(s.employess[1].FirstName, employees[1].FirstName)
				s.Equal(s.employess[1].LastName, employees[1].LastName)
				s.Equal(s.employess[1].BirthDate.Year(), employees[1].BirthDate.Year())
				s.Equal(s.employess[1].Department.Id, employees[1].Department.Id)
				s.Equal(s.employess[1].Department.Name, employees[1].Department.Name)
			},
		}, {
			desc:    "success get employees with offset and limit",
			wantErr: false,
			args: args{
				offset: 1,
				limit:  1,
			},
			checkFunc: func(args args, employees []domain.Employee) {
				s.Len(employees, 1)
				s.Equal(s.employess[1].EmployeeNumber, employees[0].EmployeeNumber)
				s.Equal(s.employess[1].FirstName, employees[0].FirstName)
				s.Equal(s.employess[1].LastName, employees[0].LastName)
				s.Equal(s.employess[1].BirthDate.Year(), employees[0].BirthDate.Year())
				s.Equal(s.employess[1].Department.Id, employees[0].Department.Id)
				s.Equal(s.employess[1].Department.Name, employees[0].Department.Name)
			},
		}, {
			desc:    "success get employees by department",
			wantErr: false,
			args: args{
				offset:  0,
				limit:   100,
				deptsId: []int{1},
			},
			checkFunc: func(args args, employees []domain.Employee) {
				s.Len(employees, 1)
				s.Equal(s.employess[0].EmployeeNumber, employees[0].EmployeeNumber)
				s.Equal(s.employess[0].FirstName, employees[0].FirstName)
				s.Equal(s.employess[0].LastName, employees[0].LastName)
				s.Equal(s.employess[0].BirthDate.Year(), employees[0].BirthDate.Year())
				s.Equal(s.employess[0].Department.Id, employees[0].Department.Id)
				s.Equal(s.employess[0].Department.Name, employees[0].Department.Name)
			},
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			employees, err := s.repo.GetEmployees(tC.args.offset, tC.args.limit, tC.args.deptsId...)
			if tC.wantErr {
				s.Error(err)
				return
			}

			s.NoError(err)
			tC.checkFunc(tC.args, employees)
		})
	}
}

func (s *EmployeeSuite) TestEmployeeRepo_InsertEmployee() {
	type (
		args struct {
			employee domain.Employee
		}
	)

	testCases := []struct {
		desc      string
		args      args
		wantErr   bool
		checkFunc func(args args, employee domain.Employee)
	}{
		{
			desc:    "success insert employee",
			wantErr: false,
			args: args{
				employee: domain.Employee{
					EmployeeNumber: "444",
					FirstName:      "Jackxx",
					LastName:       "Doexx",
					BirthDate:      time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC),
					Department: domain.Department{
						Id:   1,
						Name: "IT",
					},
				},
			},
			checkFunc: func(args args, employee domain.Employee) {
				s.NotZero(employee.Id)
				var temp Employee

				err := s.repo.GormDB.Where("id = ?", employee.Id).First(&temp).Error
				s.NoError(err)

				s.Equal(args.employee.EmployeeNumber, temp.EmployeeNumber)
				s.Equal(args.employee.FirstName, temp.FirstName)
				s.Equal(args.employee.LastName, temp.LastName)
				s.Equal(args.employee.BirthDate.Year(), temp.BirthDate.Year())
				s.Equal(args.employee.Department.Id, temp.DepartmentId)
				s.False(temp.IsDeleted)
				s.False(temp.CreatedAt.IsZero())
				s.False(temp.UpdatedAt.IsZero())
				s.Equal(temp.CreatedAt, temp.UpdatedAt)
			},
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {

			employee, err := s.repo.InsertEmployee(tC.args.employee)
			if tC.wantErr {
				s.Error(err)
				return
			}

			s.NoError(err)
			tC.checkFunc(tC.args, employee)
		})
	}

}

func (s *EmployeeSuite) TestEmployeeRepo_UpdateEmployee() {
	type (
		args struct {
			employee domain.Employee
		}
	)

	testCases := []struct {
		desc      string
		args      args
		wantErr   bool
		checkFunc func(args args, employee domain.Employee)
	}{
		{
			desc:    "success update employee",
			wantErr: false,
			args: args{
				employee: domain.Employee{
					Id:             1,
					EmployeeNumber: "111",
					FirstName:      "Johnny",
					LastName:       "Does",
					BirthDate:      time.Date(1983, 1, 1, 0, 0, 0, 0, time.UTC),
					Department: domain.Department{
						Id:   1,
						Name: "IT",
					},
				},
			},
			checkFunc: func(args args, employee domain.Employee) {
				var temp Employee

				err := s.repo.GormDB.Where("id = ?", employee.Id).First(&temp).Error
				s.NoError(err)

				s.Equal(args.employee.EmployeeNumber, temp.EmployeeNumber)
				s.Equal(args.employee.FirstName, temp.FirstName)
				s.Equal(args.employee.LastName, temp.LastName)
				s.Equal(args.employee.BirthDate.Year(), temp.BirthDate.Year())
				s.Equal(args.employee.Department.Id, temp.DepartmentId)
				s.False(temp.IsDeleted)
				s.False(temp.CreatedAt.IsZero())
				s.False(temp.UpdatedAt.IsZero())
				s.NotEqual(temp.CreatedAt, temp.UpdatedAt)
			},
		}, {
			desc:    "failed update employee duplicate employee number",
			wantErr: true,
			args: args{
				employee: domain.Employee{
					Id:             1,
					EmployeeNumber: "222",
					FirstName:      "Johnny",
					LastName:       "Does",
					BirthDate:      time.Date(1983, 1, 1, 0, 0, 0, 0, time.UTC),
					Department: domain.Department{
						Id:   1,
						Name: "IT",
					},
				},
			},
			checkFunc: func(args args, employee domain.Employee) {},
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			employee, err := s.repo.UpdateEmployee(tC.args.employee)
			if tC.wantErr {
				s.Error(err)
				return
			}

			s.NoError(err)
			tC.checkFunc(tC.args, employee)
		})
	}
}

func (s *EmployeeSuite) TestEmployeeRepo_DeleteEmployee() {
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
			desc:    "success delete employee",
			wantErr: false,
			args: args{
				id: 1,
			},
			checkFunc: func(args args) {
				var temp Employee
				err := s.repo.GormDB.Where("id = ?", args.id).First(&temp).Error
				s.NoError(err)
				s.True(temp.IsDeleted)
			},
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.repo.DeleteEmployee(tC.args.id)
			if tC.wantErr {
				s.Error(err)
				return
			}

			s.NoError(err)
			tC.checkFunc(tC.args)
		})
	}
}
