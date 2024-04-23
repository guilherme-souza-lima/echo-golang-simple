package handler

import (
	"net/http"
	"projectCRUDecho/internal/service"

	"github.com/labstack/echo/v4"
)

type UserHandlerIntercade interface {
	GetUser(c echo.Context) error
}

type UserHandler struct {
	UserServiceInterface service.UserServiceInterface
}

func NewUserHandler(UserServiceInterface service.UserServiceInterface) UserHandler {
	return UserHandler{UserServiceInterface}
}

// GetUser godoc
//
//	@Summary		GetUser.
//	@Description	GetUser.
//	@Tags			user
//	@Accept			*/*
//	@Produce		json
//	@Success		200
//	@Failure      	400
//	@Router			/ [get]
func (u UserHandler) GetUser(c echo.Context) error {
	result, err := u.UserServiceInterface.GetUSer()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
