package main

import (
	"api/config"
	"api/router"
	"api/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		// Can't load .env file
		// Continue without it
		log.Println("Can't load .env file")
	}

	config := config.LoadConfig()
	server := server.NewServer(&config)

	router.ConfigureRoutes(server)
	server.Start()
}
