package proc

import (
	"fmt"

	"github.com/kgmedia-data/gaia/pkg/msg"
	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
	"github.com/kgmedia-data/gaia/sample/internal/app/service"
	"github.com/sirupsen/logrus"
)

type EmployeeProcessor struct {
	emplSvc service.EmployeeService
}

func NewEmployeeProcessor(emplService service.EmployeeService) *EmployeeProcessor {
	return &EmployeeProcessor{emplSvc: emplService}
}

func (p EmployeeProcessor) Error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("EmployeeProcessor.(%v)(%v) %w", method, params, err)
}

func (p EmployeeProcessor) Execute(message msg.Message[domain.Employee]) error {
	employee := message.Data
	deptID := employee.Department.Id

	logrus.Infoln("EmployeeProcessor.Execute", employee.EmployeeNumber)
	if _, err := p.emplSvc.InsertEmployee(deptID, employee); err != nil {
		return err
	}

	return nil
}
