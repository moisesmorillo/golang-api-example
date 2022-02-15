package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/moisesmorillo/golang-apli-example/enums"
	"github.com/moisesmorillo/golang-apli-example/handlers"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router struct {
	server *echo.Echo
}

func NewRouter(server *echo.Echo) *Router {
	return &Router{
		server,
	}
}

func (r *Router) Init() {
	r.server.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n"}),
	)

	r.server.Use(middleware.Recover())

	basePath := r.server.Group(enums.RouterGlobalPath)
	basePath.GET(enums.RouterHealthCheckPath, handlers.HealthCheck)
	basePath.GET(enums.RouterSwaggerPath, echoSwagger.WrapHandler)
}
