package service

import "projectCRUDecho/internal/response"

type UserServiceInterface interface {
	GetUSer() (response.DtoUser, error)
}

type UserService struct{}

func NewUSerService() UserService {
	return UserService{}
}

func (u UserService) GetUSer() (response.DtoUser, error) {
	user := response.DtoUser{
		Name: "Guilherme",
		Age:  33,
	}

	return user, nil
}
