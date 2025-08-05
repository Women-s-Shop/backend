package services

import (
	"PracticalProject/models"
	"PracticalProject/repositories"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetByIdProduct(id string) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(id string, updated models.Product) (models.Product, error)
	DeleteProduct(id string) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) GetByIdProduct(id string) (models.Product, error) {
	return s.repo.GetById(id)
}

func (s *productService) CreateProduct(product models.Product) (models.Product, error) {
	return s.repo.Create(product)
}

func (s *productService) UpdateProduct(id string, updated models.Product) (models.Product, error) {
	return s.repo.Update(id, updated)
}
func (s *productService) DeleteProduct(id string) error {
	return s.repo.Delete(id)
}
