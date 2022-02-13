package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var app Application

type Application struct {
	ServerConfig *serverConfig
}

type serverConfig struct {
	Protocol string `required:"true" default:"tcp"`
	Port     string `required:"true" default:"8080"`
}

func MustLoad() Application {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading env variables:", err)
	}
	app.ServerConfig = &serverConfig{
		Protocol: os.Getenv("SERVER_PROTOCOL"),
		Port:     os.Getenv("SERVER_PORT"),
	}

	return app
}
