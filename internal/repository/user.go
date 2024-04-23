package repository

import "projectCRUDecho/entities"

type UserRepositoryInterface interface {
	InfoUserExample() (entities.DtoUser, error)
}

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (u UserRepository) InfoUserExample() (entities.DtoUser, error) {
	info := entities.DtoUser{
		Name: "Guilherme",
		Age:  33,
	}

	return info, nil
}
