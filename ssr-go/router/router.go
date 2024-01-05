package router

import (
	"ssr-go/handlers"
	s "ssr-go/server"

	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *s.Server) {
	homeHandler := handlers.NewHomeHandler()

	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())

	r := server.Echo.Group("")
	r.Static("/public", "public")
	r.GET("/", homeHandler.Handle)

}
