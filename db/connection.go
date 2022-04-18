package db

import (
	"fmt"
	"github.com/moisesmorillo/golang-api-example/config"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"
)

var (
	instance *pg.DB
	once     sync.Once
)

func GenerateClient() *pg.DB {
	once.Do(func() {
		instance = getConnection()
	})

	return instance
}

func getConnection() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Get().DbHost, config.Get().DbPort),
		User:         config.Get().DbUser,
		Password:     config.Get().DbPassword,
		Database:     config.Get().DbName,
		DialTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
	})

	db.AddQueryHook(pgdebug.DebugHook{Verbose: true})

	if _, err := db.QueryOne(pg.Discard, "select 1"); err != nil {
		log.Panicf("error generating connection to postgres: %s", err.Error())
	}

	log.Info("successfully connected to postgres")
	return db
}
