package rest

import (
	"github.com/kgmedia-data/gaia/pkg/handler"
	v1 "github.com/kgmedia-data/gaia/sample/internal/app/rest/v1"
	"github.com/kgmedia-data/gaia/sample/internal/app/service"

	echo "github.com/labstack/echo/v4"
)

type RestHandler struct {
	*handler.ServerHandler
	deptSvc     service.IDepartmentService
	employeeSvc service.IEmployeeService
}

func NewRest(host string, deptSvc service.IDepartmentService,
	employeeSvc service.IEmployeeService) *RestHandler {

	server := handler.NewServerHandler(host)

	return &RestHandler{
		ServerHandler: server,
		deptSvc:       deptSvc,
		employeeSvc:   employeeSvc,
	}
}

func (h *RestHandler) Start() error {
	h.Echo = echo.New()
	h.Echo.HideBanner = true
	h.Route()
	h.StartServer()

	return nil
}

func (h *RestHandler) Route() {
	h.ServerHandler.Route()

	v1Route := h.Echo.Group("/v1")
	v1Route.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	v1.Route(v1Route, h.deptSvc, h.employeeSvc)
}
