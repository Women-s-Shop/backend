package repositories

import (
	"PracticalProject/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrders() ([]models.Order, error)
	GetByID(id string) (models.Order, error)
	Create(order *models.Order) error
	Update(id string, fields map[string]interface{}) error
	Delete(id string) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetByID(id string) (models.Order, error) {
	var order models.Order
	err := r.db.First(&order, id).Error
	return order, err
}

func (r *orderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) Update(id string, fields map[string]interface{}) error {
	return r.db.Model(&models.Order{}).Where("id = ?", id).Updates(fields).Error
}

//	func (r *orderRepository) Update(id string, updated *models.Order) error {
//		var order models.Order
//		if err := r.db.First(&order, id).Error; err != nil {
//			return err
//		}
//		err := r.db.Model(&order).Updates(updated).Error
//		return err
//
// }
func (r *orderRepository) Delete(id string) error {
	var order models.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return err
	}
	err := r.db.Delete(&order).Error
	return err
}
