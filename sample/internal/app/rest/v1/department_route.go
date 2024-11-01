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

// GetDepartments godoc
//
//	@Summary		Get all departments
//	@Description	Get all departments
//	@Tags			department
//	@Produce		json
//	@Param			page	query	int	false	"Page"
//	@Param			limit	query	int	false	"Limit"
//	@Success		200	{object}	ResponseDoc{data=DataDoc{departments=[]departmentRespDto}}	"Success"
//	@Router			/department [get]
func (r *DepartmentRoute) getDepartments(c echo.Context) error {
	offset, limit := queryParamToOffsetLimit(c, true)
	departments, err := r.deptSvc.GetDepartments(offset, limit)
	if err != nil {
		resp := NewResponsesDto[departmentRespDto](err.Error(), nil, "departments")
		return c.JSON(500, resp)
	}

	resp := NewResponsesDto(MSG_SUCCESS, NewDepartmentsRespDto(departments), "departments")
	return c.JSON(200, resp)
}

// GetDepartment godoc
//
//	@Summary		Get department by id
//	@Description	Get department by id
//	@Tags			department
//	@Produce		json
//	@Param			deptId	path	int	true	"Department ID"
//	@Success		200	{object}	ResponseDoc{data=DataDoc{department=departmentRespDto}}	"Success"
//	@Router			/department/{deptId} [get]
func (r *DepartmentRoute) getDepartment(c echo.Context) error {
	deptId := getDepartmentId(c)
	department, err := r.deptSvc.GetDepartment(deptId)
	if err != nil {
		resp := NewResponseDto(err.Error(), nil, "department")
		return c.JSON(500, resp)
	}

	resp := NewResponseDto(MSG_SUCCESS, NewDepartmentRespDto(department), "department")
	return c.JSON(200, resp)
}

// InsertDepartment godoc
//
//	@Summary		Create new department
//	@Description	Create new department
//	@Tags			department
//	@Accept			json
//	@Produce		json
//	@Param			department	body		departmentReqDto	true	"Department data"
//	@Success		200	{object}	ResponseDoc{data=DataDoc{department=departmentRespDto}}	"Success"
//	@Router			/department [post]
func (r *DepartmentRoute) insertDepartment(c echo.Context) error {
	var deptDto departmentReqDto
	if err := c.Bind(&deptDto); err != nil {
		resp := NewResponseDto(err.Error(), nil, "department")
		return c.JSON(400, resp)
	}

	department, err := r.deptSvc.InsertDepartment(deptDto.toDomain())
	if err != nil {
		resp := NewResponseDto(err.Error(), nil, "department")
		return c.JSON(500, resp)
	}

	resp := NewResponseDto(MSG_SUCCESS, NewDepartmentRespDto(department), "department")
	return c.JSON(200, resp)
}

// UpdateDepartment godoc
//
//	@Summary		Update department by id
//	@Description	Update department by id
//	@Tags			department
//	@Accept			json
//	@Produce		json
//	@Param			deptId	path	int	true	"Department ID"
//	@Param			department	body	departmentReqDto	true	"Department data"
//	@Success		200	{object}	ResponseDoc{data=DataDoc{department=departmentRespDto}}	"Success"
//	@Router			/department/{deptId} [put]
func (r *DepartmentRoute) updateDepartment(c echo.Context) error {
	var deptDto departmentReqDto
	if err := c.Bind(&deptDto); err != nil {
		resp := NewResponseDto(err.Error(), nil, "department")
		return c.JSON(400, resp)
	}
	department := deptDto.toDomain()
	department.Id = getDepartmentId(c)

	department, err := r.deptSvc.UpdateDepartment(department)
	if err != nil {
		resp := NewResponseDto(err.Error(), nil, "department")
		return c.JSON(500, resp)
	}

	resp := NewResponseDto(MSG_SUCCESS, NewDepartmentRespDto(department), "department")
	return c.JSON(200, resp)
}

// DeleteDepartment godoc
//
//	@Summary		Delete department by id
//	@Description	Delete department by id
//	@Tags			department
//	@Produce		json
//	@Param			deptId	path	int	true	"Department ID"
//	@Success		200	{object}	ResponseDoc{data=DataDoc{department=departmentRespDto}}	"Success"
//	@Router			/department/{deptId} [delete]
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
