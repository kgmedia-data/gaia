package domain

import "time"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	BirthDate time.Time
}

func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

func (e Employee) Age() int {
	return time.Now().Year() - e.BirthDate.Year()
}

func (e Employee) IsValid() bool {
	return e.Id > 0 && e.FirstName != "" && e.LastName != ""
}
