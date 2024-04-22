package main

import (
	"context"
	"projectCRUDecho/cmd"
	"projectCRUDecho/infra"
)

func main() {
	ctx := context.Background()
	container := infra.NewContainerDI()
	cmd.StartHttp(ctx, container)
}
