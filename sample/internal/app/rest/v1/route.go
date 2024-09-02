package v1

import (
	"github.com/kgmedia-data/gaia/sample/internal/app/service"
	echo "github.com/labstack/echo/v4"
)

const (
	DEPT_ID     = "deptId"
	EMPLOYEE_ID = "employeeId"
)

const (
	MSG_SUCCESS = "success"
)

func Route(group *echo.Group, deptSvc service.IDepartmentService,
	employeeSvc service.IEmployeeService) {

}
