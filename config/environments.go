package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Environment struct {
	ServerHost string `required:"true" split_words:"true"`
	ServerPort int    `required:"true" split_words:"true"`
}

var (
	once sync.Once
	e    Environment
)

func Get() Environment {
	once.Do(func() {
		if err := envconfig.Process("", &e); err != nil {
			log.Panicf("Error parsing environment vars %#v", err)
		}
	})

	return e
}
