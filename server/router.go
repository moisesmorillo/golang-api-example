package server

import (
	"github.com/moisesmorillo/golang-api-example/enums"
	"github.com/moisesmorillo/golang-api-example/handlers"
	"github.com/moisesmorillo/golang-api-example/interfaces"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router struct {
	server       *echo.Echo
	usersHandler interfaces.UsersHandler
}

func NewRouter(server *echo.Echo, usersHandler interfaces.UsersHandler) *Router {
	return &Router{
		server,
		usersHandler,
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

	basePath.POST(enums.RouterUsersPath, r.usersHandler.Create)
	basePath.GET(enums.RouterUsersPath, r.usersHandler.Get)
}
