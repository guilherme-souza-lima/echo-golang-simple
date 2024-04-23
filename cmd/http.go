package cmd

import (
	"context"
	"projectCRUDecho/infra"

	"github.com/labstack/echo/v4"
)

func StartHttp(ctx context.Context, di *infra.ContainerDI) {
	e := echo.New()

	e.GET("/", di.UserHAndler.GetUser)

	e.Logger.Fatal(e.Start(":1323"))
}
