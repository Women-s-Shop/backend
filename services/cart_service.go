package services

import (
	"PracticalProject/models"
	"PracticalProject/repasitories"
)

type CartService interface {
	GetAll() ([]models.Cart, error)
	GetById(id uint) (*models.Cart, error)
	CreateCart(cart *models.Cart) error
	UpdateCart(cart *models.Cart) error
	DeleteCart(cart *models.Cart) error
}

type cartService struct {
	repo repasitories.CartRepository
}

func NewCartService(repo repasitories.CartRepository) CartService {
	return &cartService{repo}
}

func (s *cartService) GetAll() ([]models.Cart, error) {
	return s.repo.GetAllCarts()
}
func (s *cartService) GetById(id uint) (*models.Cart, error) {
	return s.repo.GetById(id)
}

func (s *cartService) CreateCart(cart *models.Cart) error {
	return s.repo.Create(cart)
}

func (s *cartService) UpdateCart(cart *models.Cart) error {
	return s.repo.Update(cart)
}

func (s *cartService) DeleteCart(cart *models.Cart) error {
	return s.repo.Delete(cart)
}
