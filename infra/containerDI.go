package infra

import (
	"projectCRUDecho/internal/handler"
	"projectCRUDecho/internal/service"
)

type ContainerDI struct {
	UserHAndler handler.UserHandler
	UserService service.UserService
}

func NewContainerDI() *ContainerDI {
	container := &ContainerDI{}
	container.buildService()
	container.buildHandler()
	return container
}

func (c *ContainerDI) buildService() {
	c.UserService = service.NewUSerService()
}

func (c *ContainerDI) buildHandler() {
	c.UserHAndler = handler.NewUserHandler(c.UserService)
}
