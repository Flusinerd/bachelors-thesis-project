package config

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	Port int
}

func LoadServerConfig() *ServerConfig {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	portInt, err := strconv.Atoi(port)

	if err != nil {
		panic(err)
	}

	return &ServerConfig{
		Port: portInt,
	}
}
