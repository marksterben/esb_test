package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) Routes() {
	group := s.e.Group("/api")
	group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
}
