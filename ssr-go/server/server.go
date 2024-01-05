package server

import (
	"fmt"
	"ssr-go/config"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo   *echo.Echo
	Config *config.Config
}

func NewServer() *Server {
	return &Server{
		Echo:   echo.New(),
		Config: config.LoadConfig(),
	}
}

func (server *Server) Start() {
	server.Echo.Logger.Fatal(server.Echo.Start(fmt.Sprintf(":%d", server.Config.Server.Port)))
}
