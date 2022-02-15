package server

import (
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	Container := dig.New()

	checkForInjectionError(Container.Provide(NewServer))
	checkForInjectionError(Container.Provide(NewRouter))

	return Container
}

func checkForInjectionError(err error) {
	if err != nil {
		log.Error("Error providing dependency: %s", err.Error())
	}
}
