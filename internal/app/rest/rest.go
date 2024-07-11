package rest

import (
	"github.com/kgmedia-data/gaia/internal/app/service"
	"github.com/kgmedia-data/gaia/pkg/handler"
	echo "github.com/labstack/echo/v4"
)

type RestHandler struct {
	*handler.ServerHandler
	employeeService service.IEmployeeService
}

func NewRestHandler(host string, employeeService service.IEmployeeService) *RestHandler {
	server := handler.NewServerHandler(host)
	return &RestHandler{
		ServerHandler:   server,
		employeeService: employeeService,
	}
}

func (r *RestHandler) Start() error {
	r.Echo = echo.New()
	r.Echo.HideBanner = true
	r.Route()
	r.StartServer()

	return nil
}

func (r *RestHandler) Route() {
	r.ServerHandler.Route()
	v1Route := r.Echo.Group("/v1")
	v1Route.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})
}
