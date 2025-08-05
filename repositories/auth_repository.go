package repositories

import (
	"PracticalProject/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	GetByIdEmail(email string) (*models.User, error)
}
type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}
func (r *authRepository) CreateUser(user *models.User) error {
	err := r.db.Create(user).Error
	return err
}

func (r *authRepository) GetByIdEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
