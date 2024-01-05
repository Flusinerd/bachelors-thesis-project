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
		1*time.Second,
		handler.config,
	)

	if err != nil {
		panic(err)
	}

	return c.JSON(200, products)
}