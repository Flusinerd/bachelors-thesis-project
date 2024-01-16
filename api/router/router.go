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
	server.Echo.Use(middleware.CORS())

	r := server.Echo.Group("/api")
	r.GET("/products", productsHandler.GetProducts)
	r.GET("/products/:id", productsHandler.GetProduct)
}
