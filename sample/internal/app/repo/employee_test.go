package repo

import (
	"log"
	"testing"
	"time"

	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
	"github.com/stretchr/testify/suite"
)

type EmployeeSuite struct {
	suite.Suite
	repo *EmployeeGormRepo
}

func (s *EmployeeSuite) SetupSuite() {
	db, err := NewGormRepo(TEST_CONN_STRING, 1)
	if err != nil {
		log.Fatalln(err)
	}
	truncateTables(db.GormDB, "department", "employee")

	db.GormDB.Exec("INSERT INTO department (id, name) VALUES (1, 'IT')")
	db.GormDB.Exec("INSERT INTO department (id, name) VALUES (2, 'HR')")

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

func (s *EmployeeSuite) TestEmployeeRepo() {
	employees, err := s.repo.GetEmployees(0, 100)
	s.NoError(err)
	s.Len(employees, 0)

	deptIt := domain.Department{
		Id:   1,
		Name: "IT",
	}

	deptHr := domain.Department{
		Id:   2,
		Name: "HR",
	}

	employee1 := domain.Employee{
		EmployeeNumber: "111",
		FirstName:      "John",
		LastName:       "Doe",
		BirthDate:      time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
		Department:     deptIt,
	}

	employee1, err = s.repo.InsertEmployee(employee1)
	s.NoError(err)
	s.NotZero(employee1.Id)

	employee2 := domain.Employee{
		EmployeeNumber: "222",
		FirstName:      "Jane",
		LastName:       "Doe",
		BirthDate:      time.Date(1984, 1, 1, 0, 0, 0, 0, time.UTC),
		Department:     deptHr,
	}

	employee2, err = s.repo.InsertEmployee(employee2)
	s.NoError(err)
	s.NotZero(employee2.Id)

	employees, err = s.repo.GetEmployees(0, 100)
	s.NoError(err)
	s.Len(employees, 2)
	s.Equal(employee1.EmployeeNumber, employees[0].EmployeeNumber)
	s.Equal(employee1.FirstName, employees[0].FirstName)
	s.Equal(employee1.LastName, employees[0].LastName)
	s.Equal(employee1.BirthDate.Year(), employees[0].BirthDate.Year())
	s.Equal(employee1.BirthDate.Month(), employees[0].BirthDate.Month())
	s.Equal(employee1.BirthDate.Day(), employees[0].BirthDate.Day())
	s.Equal(employee1.Department.Id, employees[0].Department.Id)
	s.Equal(employee1.Department.Name, employees[0].Department.Name)

	s.Equal(employee2.EmployeeNumber, employees[1].EmployeeNumber)
	s.Equal(employee2.FirstName, employees[1].FirstName)
	s.Equal(employee2.LastName, employees[1].LastName)
	s.Equal(employee2.BirthDate.Year(), employees[1].BirthDate.Year())
	s.Equal(employee2.BirthDate.Month(), employees[1].BirthDate.Month())
	s.Equal(employee2.BirthDate.Day(), employees[1].BirthDate.Day())
	s.Equal(employee2.Department.Id, employees[1].Department.Id)

	result, err := s.repo.GetEmployee(employee1.Id)
	s.NoError(err)
	s.Equal(employee1.EmployeeNumber, result.EmployeeNumber)
	s.Equal(employee1.FirstName, result.FirstName)
	s.Equal(employee1.LastName, result.LastName)
	s.Equal(employee1.BirthDate.Year(), result.BirthDate.Year())
	s.Equal(employee1.BirthDate.Month(), result.BirthDate.Month())
	s.Equal(employee1.BirthDate.Day(), result.BirthDate.Day())
	s.Equal(employee1.Department.Id, result.Department.Id)
	s.Equal(employee1.Department.Name, result.Department.Name)

	employee1.FirstName = "Johnny"
	employee1, err = s.repo.UpdateEmployee(employee1)
	s.NoError(err)

	result, err = s.repo.GetEmployee(employee1.Id)
	s.NoError(err)
	s.Equal(employee1.FirstName, result.FirstName)

	err = s.repo.DeleteEmployee(employee1.Id)
	s.NoError(err)

	employees, err = s.repo.GetEmployees(0, 100)
	s.NoError(err)
	s.Len(employees, 1)

	result, err = s.repo.GetEmployee(employee1.Id)
	s.Error(err)

}
