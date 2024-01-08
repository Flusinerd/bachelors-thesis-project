package product_handler

import (
	"ssr-go/client"
	"ssr-go/utils"
	"ssr-go/views/layout"
	productDetail "ssr-go/views/product-detail"

	"github.com/labstack/echo/v4"
)

func Handle(c echo.Context) error {
	productId := c.Param("id")
	if productId == "" {
		return c.String(400, "Bad Request")
	}

	product := client.GetProduct(productId)
	customerId := c.Get("customerId").(string)
	return utils.RenderView(c, layout.Base("Home", productDetail.ProductDetail(product, customerId)))
}
