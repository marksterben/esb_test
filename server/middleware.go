package server

import "github.com/labstack/echo/v4/middleware"

func (s *Server) Middlewares() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORS())
}
