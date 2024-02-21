package server

import (
	"fmt"
	"ssr-go/config"
	"ssr-go/state"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo           *echo.Echo
	Config         *config.Config
	CustomersState *state.CustomerState
}

func NewServer() *Server {
	return &Server{
		Echo:           echo.New(),
		Config:         config.LoadConfig(),
		CustomersState: state.NewCustomerState(),
	}
}

func (server *Server) Start() {
	server.Echo.Logger.Fatal(server.Echo.Start(fmt.Sprintf(":%d", server.Config.Server.Port)))
}
