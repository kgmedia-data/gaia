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

	deptRoute := NewDepartmentRoute(deptSvc)
	group.POST("/department", deptRoute.insertDepartment)
	group.GET("/department", deptRoute.getDepartments)
	group.GET("/department/:"+DEPT_ID, deptRoute.getDepartment)
	group.PUT("/department/:"+DEPT_ID, deptRoute.updateDepartment)
	group.DELETE("/department/:"+DEPT_ID, deptRoute.deleteDepartment)

	emplRoute := NewEmployeeRoute(employeeSvc)
	group.POST("/department/:"+DEPT_ID+"/employee", emplRoute.insertEmployee)
	group.GET("/department/:"+DEPT_ID+"/employee", emplRoute.getEmployees)
	group.GET("/department/:"+DEPT_ID+"/employee/:"+EMPLOYEE_ID, emplRoute.getEmployee)
	group.PUT("/department/:"+DEPT_ID+"/employee/:"+EMPLOYEE_ID, emplRoute.updateEmployee)
	group.DELETE("/department/:"+DEPT_ID+"/employee/:"+EMPLOYEE_ID, emplRoute.deleteEmployee)
}
