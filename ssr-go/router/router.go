package router

import (
	cart_handler "ssr-go/handlers/cart"
	"ssr-go/handlers/home_handler"
	"ssr-go/handlers/product_handler"
	s "ssr-go/server"

	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *s.Server) {
	homeHandler := home_handler.NewHomeHandler()
	cartHandler := cart_handler.NewCartHandler(server)
	productHandler := product_handler.NewProductHandler(server)

	server.Echo.Use(middleware.Logger())
	// server.Echo.Use(middleware.Recover())

	r := server.Echo.Group("")
	r.Static("/public", "public")
	r.GET("/", homeHandler.Handle)
	r.GET("/products/:id", productHandler.Handle)
	r.POST("/cart", cartHandler.AddToCart)
}
