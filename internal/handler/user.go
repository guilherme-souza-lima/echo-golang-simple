package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"projectCRUDecho/internal/service"
)

type UserHandler struct {
	UserServiceInterface service.UserServiceInterface
}

func NewUserHandler(UserServiceInterface service.UserServiceInterface) UserHandler {
	return UserHandler{UserServiceInterface}
}

func (u UserHandler) GetUser(c echo.Context) error {
	result, err := u.UserServiceInterface.GetUSer()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
