package handler

import (
	"context"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

type ServerHandler struct {
	host string
	Echo *echo.Echo
}

func NewServerHandler(host string) *ServerHandler {
	return &ServerHandler{
		host: host,
	}
}

func (a *ServerHandler) Start() error {
	a.Echo = echo.New()
	a.Echo.HideBanner = true
	a.Route()
	a.StartServer()

	return nil
}

func (a *ServerHandler) Route() {
	a.Echo.Use(middleware.Logger())
}

func (a *ServerHandler) StartServer() {
	go func() {
		a.Echo.Start(a.host)
	}()
}

func (a *ServerHandler) Stop() {
	a.Echo.Shutdown(context.TODO())
}
