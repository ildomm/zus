package config

import (
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/file"
	"log"
	"os"
)

var App = Configuration{}

func Setup() {
	App.Read()
}

type Configuration struct {
	Runtime  Runtime  `json:"runtime"`
	Database Database `json:"database"`
	Logger   Logger   `json:"logger"`
}

type Runtime struct {
	Project string `json:"project"`
	Port    int    `json:"port"`
	Host    string `json:"host"`
}

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Logger struct {
	BasePath   string `json:"base_path"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Compress   bool   `json:"compress"`
}

func (c *Configuration) Read() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	config.Load(file.NewSource(
		file.WithPath(dir + "/config/config.json"),
	))

	err2 := config.Scan(&c)
	log.Println(err2)

	log.Printf("Config.Runtime: %+v\n", c.Runtime)
	log.Printf("Config.Database: %+v\n", c.Database)
	log.Printf("Config.Logger: %+v\n", c.Logger)
}
