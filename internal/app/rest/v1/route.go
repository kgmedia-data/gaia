package v1

import (
	"github.com/kgmedia-data/gaia/internal/app/service"
	echo "github.com/labstack/echo/v4"
)

func Route(group *echo.Group, employeeService service.IEmployeeService) {
	employee := NewEmployeeService(employeeService)

	group.GET("/employee", employee.getEmployees)
	// group.GET("/employee/:id", employee.getEmployee)
	// group.POST("/employee", employee.insertEmployee)
	// group.PUT("/employee", employee.updateEmployee)
	// group.DELETE("/employee/:id", employee.deleteEmployee)
}
