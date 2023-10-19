package server

import (
	"esb/handler"
	"esb/repository"
	"esb/usecase"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

const (
	APP_NAME = "esb"
)

type Handler interface {
	GetCustomers(c echo.Context) error
}

type Server struct {
	e       *echo.Echo
	handler Handler
}

func NewServer() *Server {
	timeoutContext := time.Duration(viper.GetInt("TIMEOUT")) * time.Second

	repo := repository.NewRepo(DB())
	usecase := usecase.NewUsecase(repo, timeoutContext)
	handler := handler.NewHandler(usecase)

	server := &Server{
		e:       echo.New(),
		handler: handler,
	}

	return server
}

func (s *Server) Run() {
	if err := s.e.Start(viper.GetString("PORT")); err != nil {
		log.Fatal(err)
	}
}
