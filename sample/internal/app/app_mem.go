package app

import (
	"github.com/kgmedia-data/gaia/pkg/handler"
	"github.com/kgmedia-data/gaia/pkg/msg"
	"github.com/kgmedia-data/gaia/pkg/pub"
	"github.com/kgmedia-data/gaia/sample/internal/app/domain"
	"github.com/kgmedia-data/gaia/sample/internal/app/job"
	"github.com/kgmedia-data/gaia/sample/internal/app/proc"
	"github.com/kgmedia-data/gaia/sample/internal/app/repo"
	"github.com/kgmedia-data/gaia/sample/internal/app/rest"
	"github.com/kgmedia-data/gaia/sample/internal/app/service"
)

type AppMem struct {
	deptRepo *repo.DepartmentMemRepo
	emplRepo *repo.EmployeeMemRepo
	emplChan chan msg.Message[domain.Employee]
}

func NewAppMem() *AppMem {
	return &AppMem{
		deptRepo: repo.NewDepartmentMemRepo(),
		emplRepo: repo.NewEmployeeMemRepo(),
		emplChan: make(chan msg.Message[domain.Employee], 100),
	}
}

func (a *AppMem) CreateRestServer() handler.IHandler {
	deptSvc := service.NewDepartmentService(a.deptRepo)
	emplSvc := service.NewEmployeeService(a.emplRepo, a.deptRepo)

	restHhdlr := rest.NewRest(":8080", deptSvc, emplSvc)
	return restHhdlr
}

func (a *AppMem) CreateEmployeeStream() handler.IHandler {
	emplSvc := service.NewEmployeeService(a.emplRepo, a.deptRepo)
	emplProc := proc.NewEmployeeProcessor(*emplSvc)

	emplHdlr := handler.NewChanHandler[domain.Employee](a.emplChan, 1, emplProc)

	return emplHdlr
}

func (a *AppMem) CreateEmployeeJob() handler.IJob {
	emplPub := pub.NewChanPublisher(a.emplChan)
	emplJob := job.NewEmployeeJob(emplPub)

	return emplJob
}
