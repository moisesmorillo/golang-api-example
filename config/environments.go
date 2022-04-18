package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Environment struct {
	ServerHost string `required:"true" split_words:"true"`
	ServerPort int    `required:"true" split_words:"true"`
	DbHost     string `required:"true" split_words:"true"`
	DbPort     int    `required:"true" split_words:"true"`
	DbUser     string `required:"true" split_words:"true"`
	DbPassword string `required:"true" split_words:"true"`
	DbName     string `required:"true" split_words:"true"`
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
