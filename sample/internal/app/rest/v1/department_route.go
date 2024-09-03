package v1

import (
	"github.com/kgmedia-data/gaia/sample/internal/app/service"
	"github.com/labstack/echo/v4"
)

type DepartmentRoute struct {
	deptSvc service.IDepartmentService
}

func NewDepartmentRoute(deptSvc service.IDepartmentService) *DepartmentRoute {
	return &DepartmentRoute{deptSvc: deptSvc}
}

func (r *DepartmentRoute) getDepartments(c echo.Context) error {
	offset, limit := queryParamToOffsetLimit(c, true)
	departments, err := r.deptSvc.GetDepartments(offset, limit)
	if err != nil {
		resp := NewResponsesDto[departmentDto](err.Error(), nil, "departments")
		return c.JSON(500, resp)
	}

	resp := NewResponsesDto(MSG_SUCCESS, NewDepartmentsDto(departments), "departments")
	return c.JSON(200, resp)
}

func (r *DepartmentRoute) getDepartment(c echo.Context) error {
	deptId := getDepartmentId(c)
	department, err := r.deptSvc.GetDepartment(deptId)
	if err != nil {
		resp := NewResponseDto(err.Error(), nil, "department")
		return c.JSON(500, resp)
	}

	resp := NewResponseDto(MSG_SUCCESS, NewDepartmentDto(department), "department")
	return c.JSON(200, resp)
}

func (r *DepartmentRoute) insertDepartment(c echo.Context) error {
	var deptDto departmentDto
	if err := c.Bind(&deptDto); err != nil {
		resp := NewResponseDto(err.Error(), nil, "department")
		return c.JSON(400, resp)
	}

	department, err := r.deptSvc.InsertDepartment(deptDto.toDomain())
	if err != nil {
		resp := NewResponseDto(err.Error(), nil, "department")
		return c.JSON(500, resp)
	}

	resp := NewResponseDto(MSG_SUCCESS, NewDepartmentDto(department), "department")
	return c.JSON(200, resp)
}

func (r *DepartmentRoute) updateDepartment(c echo.Context) error {
	var deptDto departmentDto
	if err := c.Bind(&deptDto); err != nil {
		resp := NewResponseDto(err.Error(), nil, "department")
		return c.JSON(400, resp)
	}
	deptDto.Id = getDepartmentId(c)
	department, err := r.deptSvc.UpdateDepartment(deptDto.toDomain())
	if err != nil {
		resp := NewResponseDto(err.Error(), nil, "department")
		return c.JSON(500, resp)
	}

	resp := NewResponseDto(MSG_SUCCESS, NewDepartmentDto(department), "department")
	return c.JSON(200, resp)
}

func (r *DepartmentRoute) deleteDepartment(c echo.Context) error {
	deptId := getDepartmentId(c)
	err := r.deptSvc.DeleteDepartment(deptId)
	if err != nil {
		resp := NewResponseDto(err.Error(), nil, "department")
		return c.JSON(500, resp)
	}

	resp := NewResponseDto(MSG_SUCCESS, nil, "department")
	return c.JSON(200, resp)
}
