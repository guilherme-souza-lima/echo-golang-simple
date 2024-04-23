package infra

import (
	"projectCRUDecho/internal/handler"
	"projectCRUDecho/internal/repository"
	"projectCRUDecho/internal/service"
)

type ContainerDI struct {
	Config         Config
	UserHAndler    handler.UserHandler
	UserService    service.UserService
	UserRepository repository.UserRepository
}

func NewContainerDI(config Config) *ContainerDI {
	container := &ContainerDI{
		Config: config,
	}
	container.buildRepository()
	container.buildService()
	container.buildHandler()
	return container
}

func (c *ContainerDI) buildRepository() {
	c.UserRepository = repository.NewUserRepository()
}

func (c *ContainerDI) buildService() {
	c.UserService = service.NewUSerService(c.UserRepository)
}

func (c *ContainerDI) buildHandler() {
	c.UserHAndler = handler.NewUserHandler(c.UserService)
}
