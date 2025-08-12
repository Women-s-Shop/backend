package repositories

import (
	"PracticalProject/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB}
}

func (repo *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := repo.DB.Find(&users).Error
	return users, err
}

func (repo *UserRepository) GetByID(id string) (models.User, error) {
	var user models.User
	err := repo.DB.First(&user, id).Error
	return user, err
}

func (repo *UserRepository) Create(user models.User) error {
	return repo.DB.Create(&user).Error
}

func (repo *UserRepository) Update(user models.User) error {
	return repo.DB.Save(&user).Error
}

func (repo *UserRepository) Delete(id string) error {
	var user models.User
	if err := repo.DB.First(&user, id).Error; err != nil {
		return err
	}
	return repo.DB.Delete(&user).Error
}
