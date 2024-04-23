package main

import (
	"context"
	"projectCRUDecho/cmd"
	"projectCRUDecho/infra"
)

func main() {
	config := infra.NewConfig()
	container := infra.NewContainerDI(config)

	cmd.StartHttp(
		context.Background(),
		container,
	)
}
