package service

import (
	"fmt"
	"projectCRUDecho/infra/token"
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

	tokenMaker, err := token.NewPaseto("abcdefghijkl12345678901234567890")
	if err != nil {
		return user, fmt.Errorf("Couldn't create token maker: %w", err)
	}

	info, _ := tokenMaker.CreateToken("guilherme", time.Minute)
	fmt.Println(info)

	now := time.Now()
	user = response.ResponseUser{
		Name: result.Name,
		Age:  result.Age,
		Now:  now,
	}
	return user, nil
}
