package service

import (
	"projectCRUDecho/internal/repository"
	"projectCRUDecho/internal/response"
	"time"
)

type UserServiceInterface interface {
	GetUSer() (response.ResponseUser, error)
}

type UserService struct {
	UserRepositoryInterface repository.UserRepositoryInterface
}

func NewUSerService(UserRepositoryInterface repository.UserRepositoryInterface) UserService {
	return UserService{UserRepositoryInterface}
}

func (u UserService) GetUSer() (user response.ResponseUser, err error) {
	result, err := u.UserRepositoryInterface.InfoUserExample()
	if err != nil {
		return user, err
	}

	now := time.Now()
	user = response.ResponseUser{
		Name: result.Name,
		Age:  result.Age,
		Now:  now,
	}
	return user, nil
}
