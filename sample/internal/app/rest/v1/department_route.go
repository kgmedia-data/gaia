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
	departments, err := r.deptSvc.GetDepartments(0, 10)
	if err != nil {
		return c.JSON(500, err)
	}

	resp := NewResponseDtos(MSG_SUCCESS, NewDepartmentsDto(departments), "departments")
	return c.JSON(200, resp)
}
