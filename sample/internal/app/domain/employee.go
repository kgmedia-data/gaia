package domain

import "time"

type Employee struct {
	Id             int
	EmployeeNumber string
	FirstName      string
	LastName       string
	BirthDate      time.Time
	Department     Department
}

func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

func (e Employee) Age() int {
	return time.Now().Year() - e.BirthDate.Year()
}

func (e Employee) IsValid() bool {
	return e.EmployeeNumber != "" &&
		e.FirstName != "" &&
		e.LastName != ""
}
