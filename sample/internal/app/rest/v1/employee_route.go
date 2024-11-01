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

// GetEmployees godoc
//
//	@Summary		Get all employees
//	@Description	Get all employees
//	@Tags			employee
//	@Produce		json
//	@Param			page	query	int	false	"Page"
//	@Param			limit	query	int	false	"Limit"
//	@Param			deptId	path	int	true	"Department ID"
//	@Success		200	{object}	ResponseDoc{data=DataDoc{employees=[]employeeRespDto}}	"Success"
//	@Router			/department/{deptId}/employee [get]
func (r *EmployeeRoute) getEmployees(c echo.Context) error {
	offset, limit := queryParamToOffsetLimit(c, true)
	deptId := getDepartmentId(c)

	employees, err := r.emplSvc.GetEmployees(deptId, offset, limit)
	if err != nil {
		resp := NewResponsesDto[employeeRespDto](err.Error(), nil, "employees")
		return c.JSON(500, resp)
	}

	resp := NewResponsesDto(MSG_SUCCESS, NewEmployeesRespDto(employees), "employees")
	return c.JSON(200, resp)
}

// GetEmployee godoc
//
//	@Summary		Get employee by id
//	@Description	Get employee by id
//	@Tags			employee
//	@Produce		json
//	@Param			deptId	path	int	true	"Department ID"
//	@Param			employeeId	path	int	true	"Employee ID"
//	@Success		200	{object}	ResponseDoc{data=DataDoc{employee=employeeRespDto}}	"Success"
//	@Router			/department/{deptId}/employee/{employeeId} [get]
func (r *EmployeeRoute) getEmployee(c echo.Context) error {
	deptId := getDepartmentId(c)
	emplId := getEmployeeId(c)

	employee, err := r.emplSvc.GetEmployee(deptId, emplId)
	if err != nil {
		resp := NewResponseDto(err.Error(), nil, "employee")
		return c.JSON(500, resp)
	}

	resp := NewResponseDto(MSG_SUCCESS, NewEmployeeRespDto(employee), "employee")
	return c.JSON(200, resp)
}

// InsertEmployee godoc
//
//	@Summary		Create new employee
//	@Description	Create new employee
//	@Tags			employee
//	@Accept			json
//	@Produce		json
//	@Param			deptId	path	int	true	"Department ID"
//	@Success		200	{object}	ResponseDoc{data=DataDoc{employee=employeeRespDto}}	"Success"
//	@Router			/department/{deptId}/employee [post]
func (r *EmployeeRoute) insertEmployee(c echo.Context) error {
	var emplDto employeeReqDto
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

	resp := NewResponseDto(MSG_SUCCESS, NewEmployeeRespDto(employee), "employee")
	return c.JSON(200, resp)
}

// UpdateEmployee godoc
//
//	@Summary		Update employee by id
//	@Description	Update employee by id
//	@Tags			employee
//	@Accept			json
//	@Produce		json
//	@Param			deptId	path	int	true	"Department ID"
//	@Param			employeeId	path	int	true	"Employee ID"
//	@Param			employee	body	employeeReqDto	true	"Employee data"
//	@Success		200	{object}	ResponseDoc{data=DataDoc{employee=employeeRespDto}}	"Success"
//	@Router			/department/{deptId}/employee/{employeeId} [put]
func (r *EmployeeRoute) updateEmployee(c echo.Context) error {
	var emplDto employeeReqDto
	if err := c.Bind(&emplDto); err != nil {
		resp := NewResponseDto(err.Error(), nil, "employee")
		return c.JSON(400, resp)
	}
	deptId := getDepartmentId(c)
	empl := emplDto.toDomain()
	empl.Id = getEmployeeId(c)

	employee, err := r.emplSvc.UpdateEmployee(deptId, empl)
	if err != nil {
		resp := NewResponseDto(err.Error(), nil, "employee")
		return c.JSON(500, resp)
	}

	resp := NewResponseDto(MSG_SUCCESS, NewEmployeeRespDto(employee), "employee")
	return c.JSON(200, resp)
}

// DeleteEmployee godoc
//
//	@Summary		Delete employee by id
//	@Description	Delete employee by id
//	@Tags			employee
//	@Produce		json
//	@Param			deptId	path	int	true	"Department ID"
//	@Param			employeeId	path	int	true	"Employee ID"
//	@Success		200	{object}	ResponseDoc{data=DataDoc{employee=employeeRespDto}}	"Success"
//	@Router			/department/{deptId}/employee/{employeeId} [delete]
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
