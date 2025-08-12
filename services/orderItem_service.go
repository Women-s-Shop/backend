package services

import (
	"PracticalProject/models"
	"PracticalProject/repositories"
	"errors"
)

type OrderItemService struct {
	Repo *repositories.OrderItemRepository
}

func NewOrderItemService(repo *repositories.OrderItemRepository) *OrderItemService {
	return &OrderItemService{repo}
}

func (service *OrderItemService) GetAll() ([]models.OrderItem, error) {
	return service.Repo.GetAll()
}

func (service *OrderItemService) GetByID(id string) (models.OrderItem, error) {
	orderItem, err := service.Repo.GetByID(id)
	if err != nil {
		return models.OrderItem{}, errors.New("OrderItem not found")
	}
	return orderItem, nil
}

func (service *OrderItemService) Create(orderItem models.OrderItem) (models.OrderItem, error) {
	err := service.Repo.Create(orderItem)
	if err != nil {
		return models.OrderItem{}, err
	}
	return orderItem, nil
}

func (service *OrderItemService) Update(orderItem models.OrderItem) (models.OrderItem, error) {
	err := service.Repo.Update(orderItem)
	if err != nil {
		return models.OrderItem{}, err
	}
	return orderItem, nil
}

func (service *OrderItemService) Delete(id string) error {
	err := service.Repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
