package server

import (
	"context"
	"log"

	"github.com/F7icK/pc-accessories-store/internal/server/handlers"
	"github.com/F7icK/pc-accessories-store/internal/types/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	httpServer *echo.Echo
	portHTTP   string
}

func NewServer(cfg *config.Config, handlers *handlers.Handlers) *Server {
	router := NewRouter(handlers)

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	return &Server{
		httpServer: router,
		portHTTP:   cfg.HTTP.Port,
	}
}

func (s *Server) Run() error {
	log.Println("Restart server!")
	if err := s.httpServer.Start(s.portHTTP); err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
