package main

import (
	"projectCRUDecho/infra"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title Swagger Echo-Golang-Simples API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhot:1323
// @BasePath /
func main() {
	config := infra.NewConfig()
	container := infra.NewContainerDI(config)

	container.Server.Start()
}
