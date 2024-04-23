package main

import (
	"projectCRUDecho/infra"
)

func main() {
	config := infra.NewConfig()
	container := infra.NewContainerDI(config)

	container.Server.Start()
}
