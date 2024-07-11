package v1

import (
	"github.com/kgmedia-data/gaia/internal/app/service"
	"github.com/labstack/echo/v4"
)

type EmployeeRoute struct {
	employeeService service.IEmployeeService
}

func NewEmployeeService(employeeService service.IEmployeeService) *EmployeeRoute {
	return &EmployeeRoute{
		employeeService: employeeService,
	}
}

func (e *EmployeeRoute) getEmployees(c echo.Context) error {
	return nil
}
