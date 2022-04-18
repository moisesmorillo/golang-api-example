package main

import (
	"fmt"
	"github.com/moisesmorillo/golang-api-example/config"
	"github.com/moisesmorillo/golang-api-example/server"

	_ "github.com/moisesmorillo/golang-api-example/docs"

	"github.com/labstack/echo/v4"
)

// @title API Example
// @version 1.0
// @description show echo capabilities

// @license.name Platzi Live Class

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
