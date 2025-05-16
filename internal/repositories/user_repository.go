package repositories

import (
	"point-of-sales-app/internal/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]entities.User, error)
	GetUser(id int) (*entities.User, error)
	UpdateUser(id int, updateData map[string]interface{}) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUsers() ([]entities.User, error) {
	var users []entities.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) GetUser(id int) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) UpdateUser(id int, updateData map[string]interface{}) error {
	return r.db.Model(&entities.User{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *userRepository) DeleteUser(id int) error {
	return r.db.Delete(&entities.User{}, id).Error
}
