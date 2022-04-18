package server

import (
	"github.com/moisesmorillo/golang-api-example/db"
	"github.com/moisesmorillo/golang-api-example/db/repositories"
	"github.com/moisesmorillo/golang-api-example/handlers"
	"github.com/moisesmorillo/golang-api-example/services"

	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	Container := dig.New()

	checkForInjectionError(Container.Provide(db.GenerateClient))
	checkForInjectionError(Container.Provide(repositories.NewUsersRepository))

	checkForInjectionError(Container.Provide(services.NewUsersService))

	checkForInjectionError(Container.Provide(handlers.NewUsersHandler))

	checkForInjectionError(Container.Provide(NewServer))
	checkForInjectionError(Container.Provide(NewRouter))

	return Container
}

func checkForInjectionError(err error) {
	if err != nil {
		log.Errorf("Error providing dependency: %s", err.Error())
	}
}
