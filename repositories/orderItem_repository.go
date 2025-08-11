package repositories

import (
	"PracticalProject/models"
	"gorm.io/gorm"
)

type OrderItemRepository struct {
	DB *gorm.DB
}

func (repo *OrderItemRepository) GetAll() ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := repo.DB.Find(&orderItems).Error
	return orderItems, err
}

func (repo *OrderItemRepository) GetByID(id string) (models.OrderItem, error) {
	var orderItem models.OrderItem
	err := repo.DB.First(&orderItem, id).Error
	return orderItem, err
}

func (repo *OrderItemRepository) Create(orderItem models.OrderItem) error {
	return repo.DB.Create(&orderItem).Error
}

func (repo *OrderItemRepository) Update(orderItem models.OrderItem) error {
	return repo.DB.Save(&orderItem).Error
}

func (repo *OrderItemRepository) Delete(id string) error {
	var orderItem models.OrderItem
	if err := repo.DB.First(&orderItem, id).Error; err != nil {
		return err
	}
	return repo.DB.Delete(&orderItem).Error
}
