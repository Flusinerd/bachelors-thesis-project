package main

import (
	"ssr-go/middleware"
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
	server.Echo.Use(middleware.CustomerIdMiddleware)
	router.ConfigureRoutes(server)

	server.Start()
}
