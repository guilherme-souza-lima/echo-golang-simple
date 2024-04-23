package cmd

import (
	"projectCRUDecho/internal/handler"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	UserHandlerIntercade handler.UserHandlerIntercade
}

func NewServer(UserHandlerIntercade handler.UserHandlerIntercade) Server {
	return Server{UserHandlerIntercade}
}

func (s Server) Start() {
	e := echo.New()

	e.GET("/", s.UserHandlerIntercade.GetUser)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
