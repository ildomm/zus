package db

import (
	"fmt"
	"github.com/ildomm/zus/config"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
)

/* Return Postgres session instance */
func Postgres() *dbr.Session {
	return conn.Postgres.NewSession(nil)
}
func postgres() *dbr.Connection {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.App.Database.Host,
		config.App.Database.Port,
		config.App.Database.Username,
		config.App.Database.Password,
		config.App.Database.Database)

	log.Println("PostgreSQL: " + psqlInfo)

	/* See: https://github.com/gocraft/dbr
	wiki -> https://godoc.org/github.com/gocraft/dbr
	*/
	conn, err := dbr.Open("postgres", psqlInfo, nil )
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(4)

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return conn
}

func ResetPostgresDB() bool {

	// Postgres
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.App.Database.Host,
		config.App.Database.Port,
		config.App.Database.Username,
		config.App.Database.Password,
		config.App.Database.Database)

	// Postgres: Drop
	_, err := exec.Command("migrate",
		"-source",
		"file://database/migrations",
		"-database",
		psqlInfo,
		"drop").Output()
	if err != nil {
		log.Fatal(err)
		return false
	}

	// Postgres: UP
	files, err := ioutil.ReadDir("./database/migrations")
	if err != nil {
		log.Fatal(err)
		return false
	}

	var lastFile string
	for _, f := range files {
		lastFile = f.Name()
	}
	lastMigration := strings.Split(lastFile, "_")[0]

	_, err = exec.Command("migrate",
		"-source",
		"file://database/migrations",
		"-database",
		psqlInfo,
		"goto", lastMigration).Output()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func CleanPostgresDB() bool {
	tablesTruncate := []string{
		"tokens",
	}

	for _, table := range tablesTruncate {
		_, err := Postgres().Exec("DELETE FROM " + table)
		if err != nil {
			log.Fatal(err)
			return false
		}
	}

	return true
}

func CountTable(table string) int {
	var count int = 0
	session := Postgres()
	session.Select("COUNT(*) as count").From(table).
		LoadOne(&count)

	return count
}