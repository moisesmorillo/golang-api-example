package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/moisesmorillo/golang-apli-example/config"
	_ "github.com/moisesmorillo/golang-apli-example/docs"
	"github.com/moisesmorillo/golang-apli-example/server"
)

// @title API Example
// @version 1.0
// @description show echo capabilities

// @BasePath /api/example
// @schemes http https
func main() {
	containers := server.BuildContainer()

	if err := containers.Invoke(func(router *server.Router, server *echo.Echo) {
		router.Init()

		if servErr := server.Start(fmt.Sprintf("%s:%d", config.Get().ServerHost, config.Get().ServerPort)); servErr != nil {
			panic(servErr)
		}
	}); err != nil {
		panic(err)
	}

}
