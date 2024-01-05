package handlers

import (
	"ssr-go/utils"
	views "ssr-go/views/home"
	"ssr-go/views/layout"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Handle(c echo.Context) error {
	return utils.RenderView(c, layout.Base("Home", views.Home()))
}
