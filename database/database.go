package db

import (
	"log"
	"sync"

	"github.com/gocraft/dbr"
)

var conn *DATABASES
var once sync.Once

type DATABASES struct {
	Postgres *dbr.Connection
}

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
)

func Setup() *DATABASES {
	once.Do(func() {
		conn = instances()
	})
	return conn
}

func instances() *DATABASES {
	log.Println("Connecting databases")

	return &DATABASES{
		Postgres: postgres(),
	}
}

func CleanDatabases() {
	CleanPostgresDB()
}
