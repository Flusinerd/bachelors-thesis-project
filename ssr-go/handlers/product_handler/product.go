package product_handler

import (
	"ssr-go/client"
	"ssr-go/components/alert"
	s "ssr-go/server"
	"ssr-go/utils"
	"ssr-go/views/layout"
	productDetail "ssr-go/views/product-detail"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	server *s.Server
}

func NewProductHandler(server *s.Server) *ProductHandler {
	return &ProductHandler{
		server: server,
	}
}

func (h *ProductHandler) Handle(c echo.Context) error {
	productId := c.Param("id")
	if productId == "" {
		return c.String(400, "Bad Request")
	}

	product := client.GetProduct(productId)
	customerId := c.Get("customerId").(string)
	added := c.QueryParam("added")
	var alertCmp templ.Component = nil
	if added == "true" {
		alertCmp = alert.AlertComponent("Product added to cart", "success")
	}

	customerCart := h.server.CustomersState.GetCart(customerId)

	return utils.RenderView(c, layout.Base("Home", productDetail.ProductDetail(product, customerId), alertCmp, &customerCart))
}
