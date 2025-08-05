package repasitories

import (
	"PracticalProject/models"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetAllCarts() ([]models.Cart, error)
	GetById(id uint) (*models.Cart, error)
	Create(cart *models.Cart) error
	Update(cart *models.Cart) error
	Delete(cart *models.Cart) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}
func (r *cartRepository) GetAllCarts() ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Find(&carts).Error
	return carts, err
}

func (r *cartRepository) GetById(id uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.First(&cart, id).Error
	return &cart, err
}

func (r *cartRepository) Create(cart *models.Cart) error {
	err := r.db.Create(cart).Error
	return err
}

func (r *cartRepository) Update(cart *models.Cart) error {
	err := r.db.Save(cart).Error
	return err
}

func (r *cartRepository) Delete(cart *models.Cart) error {
	err := r.db.Delete(cart).Error
	return err
}
