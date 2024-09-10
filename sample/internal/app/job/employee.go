package job

import (
	"fmt"
	"time"

	"github.com/kgmedia-data/gaia/pkg/msg"
	"github.com/kgmedia-data/gaia/pkg/pub"
	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
	"github.com/sirupsen/logrus"
)

type EmployeeJob struct {
	emplPub pub.IPublisher[domain.Employee]
}

func NewEmployeeJob(emplPub pub.IPublisher[domain.Employee]) *EmployeeJob {
	return &EmployeeJob{emplPub: emplPub}
}

func (j EmployeeJob) error(err error, method string, params ...interface{}) error {
	return fmt.Errorf("EmployeeJob.(%v)(%v) %w", method, params, err)
}

func (j EmployeeJob) Run() error {
	employees := j.generateEmployee(10)

	for _, employee := range employees {
		data := msg.Message[domain.Employee]{Data: employee}
		logrus.Infoln("EmployeeJob.Run", employee.EmployeeNumber)
		if err := j.emplPub.Publish(data); err != nil {
			logrus.Errorln(j.error(err, "Run"))
		}
	}

	return nil
}

func (j EmployeeJob) generateEmployee(n int) []domain.Employee {
	var result []domain.Employee

	for i := 0; i < n; i++ {
		result = append(result, domain.Employee{
			EmployeeNumber: fmt.Sprintf("EMP%03d", i),
			FirstName:      fmt.Sprintf("First%03d", i),
			LastName:       fmt.Sprintf("Last%03d", i),
			BirthDate:      time.Now().AddDate(-i*5, -i, -i),
			Department:     domain.Department{Id: 1, Name: "IT"},
		})
	}

	return result
	// generate employee
}
