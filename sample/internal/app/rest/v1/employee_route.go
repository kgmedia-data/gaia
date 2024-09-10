package v1

import (
	"github.com/kgmedia-data/gaia/sample/internal/app/service"
	"github.com/labstack/echo/v4"
)

type EmployeeRoute struct {
	emplSvc service.IEmployeeService
}

func NewEmployeeRoute(emplSvc service.IEmployeeService) *EmployeeRoute {
	return &EmployeeRoute{emplSvc: emplSvc}
}

func (r *EmployeeRoute) getEmployees(c echo.Context) error {
	offset, limit := queryParamToOffsetLimit(c, true)
	deptId := getDepartmentId(c)

	employees, err := r.emplSvc.GetEmployees(deptId, offset, limit)
	if err != nil {
		resp := NewResponsesDto[employeeDto](err.Error(), nil, "employees")
		return c.JSON(500, resp)
	}

	resp := NewResponsesDto(MSG_SUCCESS, NewEmployeesDto(employees), "employees")
	return c.JSON(200, resp)
}

func (r *EmployeeRoute) getEmployee(c echo.Context) error {
	deptId := getDepartmentId(c)
	emplId := getEmployeeId(c)

	employee, err := r.emplSvc.GetEmployee(deptId, emplId)
	if err != nil {
		resp := NewResponseDto(err.Error(), nil, "employee")
		return c.JSON(500, resp)
	}

	resp := NewResponseDto(MSG_SUCCESS, NewEmployeeDto(employee), "employee")
	return c.JSON(200, resp)
}

func (r *EmployeeRoute) insertEmployee(c echo.Context) error {
	var emplDto employeeDto
	if err := c.Bind(&emplDto); err != nil {
		resp := NewResponseDto(err.Error(), nil, "employee")
		return c.JSON(400, resp)
	}
	deptId := getDepartmentId(c)

	employee, err := r.emplSvc.InsertEmployee(deptId, emplDto.toDomain())
	if err != nil {
		resp := NewResponseDto(err.Error(), nil, "employee")
		return c.JSON(500, resp)
	}

	resp := NewResponseDto(MSG_SUCCESS, NewEmployeeDto(employee), "employee")
	return c.JSON(200, resp)
}

func (r *EmployeeRoute) updateEmployee(c echo.Context) error {
	var emplDto employeeDto
	if err := c.Bind(&emplDto); err != nil {
		resp := NewResponseDto(err.Error(), nil, "employee")
		return c.JSON(400, resp)
	}
	deptId := getDepartmentId(c)
	emplDto.Id = getEmployeeId(c)

	employee, err := r.emplSvc.UpdateEmployee(deptId, emplDto.toDomain())
	if err != nil {
		resp := NewResponseDto(err.Error(), nil, "employee")
		return c.JSON(500, resp)
	}

	resp := NewResponseDto(MSG_SUCCESS, NewEmployeeDto(employee), "employee")
	return c.JSON(200, resp)
}

func (r *EmployeeRoute) deleteEmployee(c echo.Context) error {
	deptId := getDepartmentId(c)
	emplId := getEmployeeId(c)

	err := r.emplSvc.DeleteEmployee(deptId, emplId)
	if err != nil {
		resp := NewResponseDto(err.Error(), nil, "employee")
		return c.JSON(500, resp)
	}

	resp := NewResponseDto(MSG_SUCCESS, nil, "employee")
	return c.JSON(200, resp)
}
