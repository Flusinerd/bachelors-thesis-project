package cart_handler

import (
	"ssr-go/client"
	"ssr-go/server"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	server *server.Server
}

func NewCartHandler(server *server.Server) *CartHandler {
	return &CartHandler{
		server: server,
	}
}

func (h *CartHandler) AddToCart(c echo.Context) error {
	customerId := c.Get("customerId").(string)
	productId := c.FormValue("product_id")
	product := client.GetProduct(productId)

	if customerId == "" || productId == "" {
		return echo.ErrBadRequest
	}

	h.server.CustomersState.AddProduct(customerId, *product)

	return c.Redirect(302, "/products/"+productId+"?added=true")
}
