package cmd

import (
	_ "projectCRUDecho/docs"
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

//	@title			Swagger GO-ECHO API
//	@version		1.0
//	@description	Documentação API

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

// @host			localhost:1323
func (s Server) Start() {
	e := echo.New()

	e.GET("/", s.UserHandlerIntercade.GetUser)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
