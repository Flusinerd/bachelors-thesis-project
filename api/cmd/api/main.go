package main

import (
	"api/config"
	"api/router"
	"api/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config := config.LoadConfig()
	server := server.NewServer(&config)

	router.ConfigureRoutes(server)
	server.Start()
}
