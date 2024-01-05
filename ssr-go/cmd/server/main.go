package main

import (
	"ssr-go/router"
	"ssr-go/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	server := server.NewServer()
	router.ConfigureRoutes(server)

	server.Start()
}
