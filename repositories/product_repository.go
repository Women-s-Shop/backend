package repositories

import (
	"PracticalProject/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetById(id string) (models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(id string, updated models.Product) (models.Product, error)
	Delete(id string) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var product []models.Product
	err := r.db.Find(&product).Error
	return product, err
}

func (r *productRepository) GetById(id string) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *productRepository) Create(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *productRepository) Update(id string, updated models.Product) (models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return product, err
	}
	err := r.db.Model(&product).Updates(updated).Error
	return product, err
}

func (r *productRepository) Delete(id string) error {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return err
	}
	err := r.db.Delete(&product).Error
	return err
}
