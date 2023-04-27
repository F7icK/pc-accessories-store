package handlers

import (
	"net/http"

	"github.com/F7icK/pc-accessories-store/internal/service"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	s *service.Service
}

func NewHandlers(s *service.Service) *Handlers {
	return &Handlers{
		s: s,
	}
}

func (h *Handlers) Ping(c echo.Context) error {
	if err := c.String(http.StatusOK, "Pong"); err != nil {
		return err
	}
	return nil
}
