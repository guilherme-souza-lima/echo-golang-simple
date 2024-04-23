package infra

import (
	"projectCRUDecho/cmd"
	"projectCRUDecho/internal/handler"
	"projectCRUDecho/internal/repository"
	"projectCRUDecho/internal/service"
)

type ContainerDI struct {
	Config         Config
	Server         cmd.Server
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
	container.buildHttp()

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

func (c *ContainerDI) buildHttp() {
	c.Server = cmd.NewServer(c.UserHAndler)
}
