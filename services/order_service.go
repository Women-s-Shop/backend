package services

import (
	"PracticalProject/models"
	"PracticalProject/repositories"
)

type OrderService interface {
	GetAll() ([]models.Order, error)
	GetByID(id string) (models.Order, error)
	CreateOrder(order *models.Order) error
	UpdateOrder(id string, fields map[string]interface{}) error
	DeleteOrder(id string) error
}

type orderService struct {
	repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) OrderService {
	return &orderService{repo}
}

func (s *orderService) GetAll() ([]models.Order, error) {
	return s.repo.GetAllOrders()
}

func (s *orderService) GetByID(id string) (models.Order, error) {
	return s.repo.GetByID(id)
}

func (s *orderService) CreateOrder(order *models.Order) error {
	return s.repo.Create(order)
}

func (s *orderService) UpdateOrder(id string, fields map[string]interface{}) error {
	return s.repo.Update(id, fields)
}

//func (s *orderService) UpdateOrder(id string, updated *models.Order) error {
//	return s.repo.Update(id, updated)
//}

func (s *orderService) DeleteOrder(id string) error {
	return s.repo.Delete(id)
}
