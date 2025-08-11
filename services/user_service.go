package services

import (
	"PracticalProject/models"
	"PracticalProject/repositories"
	"errors"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (service *UserService) GetAll() ([]models.User, error) {
	return service.Repo.GetAll()
}

func (service *UserService) GetByID(id string) (models.User, error) {
	user, err := service.Repo.GetByID(id)
	if err != nil {
		return models.User{}, errors.New("User not found")
	}
	return user, nil
}

func (service *UserService) Create(user models.User) (models.User, error) {
	err := service.Repo.Create(user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (service *UserService) Update(user models.User) (models.User, error) {
	err := service.Repo.Update(user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (service *UserService) Delete(id string) error {
	err := service.Repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
