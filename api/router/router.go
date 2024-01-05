package router

import (
	"api/handlers"
	s "api/server"

	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *s.Server) {
	productsHandler := handlers.NewProductHandler(server)

	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())

	r := server.Echo.Group("/api")
	r.GET("/products", productsHandler.GetProducts)
}
