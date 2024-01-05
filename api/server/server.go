package server

import (
	"api/config"
	"api/repositories"
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo         *echo.Echo
	Repositories *repositories.Repositories
	Config       *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Echo:         echo.New(),
		Repositories: repositories.NewRepositories(),
		Config:       cfg,
	}
}

func (server *Server) Start() {
	server.Echo.Logger.Fatal(server.Echo.Start(fmt.Sprintf(":%d", server.Config.Server.Port)))
}
