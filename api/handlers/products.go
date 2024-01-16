package handlers

import (
	c "api/config"
	"api/repositories"
	s "api/server"
	"api/utils"
	"time"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productsRepository repositories.ProductsRepository
	config             *c.Config
}

type ProductHandler interface {
	GetProducts(c echo.Context) error
	GetProduct(c echo.Context) error
}

func NewProductHandler(server *s.Server) ProductHandler {
	return &productHandler{
		productsRepository: server.Repositories.Products,
		config:             server.Config,
	}
}

func (handler *productHandler) GetProducts(c echo.Context) error {
	products, err := utils.TimeoutWrapper(
		func() interface{} { return handler.productsRepository.GetProducts() },
		500*time.Millisecond,
		handler.config,
	)

	if err != nil {
		panic(err)
	}

	return c.JSON(200, products)
}

func (handler *productHandler) GetProduct(c echo.Context) error {
	id := c.Param("id")
	product, err := utils.TimeoutWrapper(
		func() interface{} { return handler.productsRepository.GetProductById(id) },
		100*time.Millisecond,
		handler.config,
	)

	if err != nil {
		panic(err)
	}

	return c.JSON(200, product)
}
