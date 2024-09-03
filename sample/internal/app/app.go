package app

import (
	"github.com/kgmedia-data/gaia/pkg/handler"
	"github.com/kgmedia-data/gaia/sample/config"
	"github.com/kgmedia-data/gaia/sample/internal/app/repo"
	"github.com/kgmedia-data/gaia/sample/internal/app/rest"
	"github.com/kgmedia-data/gaia/sample/internal/app/service"
	"github.com/sirupsen/logrus"
)

type App struct {
	cfg config.Config
}

func NewApp(cfg config.Config) *App {
	return &App{cfg: cfg}
}

func (a *App) CreateRestServer() handler.IHandler {
	gormRepo, err := repo.NewGormRepo(a.cfg.Rest.Db.DataStore, a.cfg.Rest.Db.NumberConn)
	if err != nil {
		logrus.Fatalln(err)
	}
	deptRepo := repo.NewDepartmentGorm(gormRepo)
	emplRepo := repo.NewEmployeeGorm(gormRepo)

	deptSvc := service.NewDepartmentService(deptRepo)
	emplSvc := service.NewEmployeeService(emplRepo, deptRepo)

	restHdlr := rest.NewRest(a.cfg.Rest.Server.Host, deptSvc, emplSvc)
	return restHdlr
}
