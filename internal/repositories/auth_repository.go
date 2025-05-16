package repositories

import (
	"log"
	"point-of-sales-app/internal/entities"

	"gorm.io/gorm"
)

type AuthRepository interface {
	EmailExists(email string) bool
	GetUserByEmail(email string) (*entities.User, error)
	CreateUser(user *entities.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(DB *gorm.DB) AuthRepository {
	return &authRepository{
		db: DB,
	}
}

func (r *authRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user *entities.User

	err := r.db.Take(&user, "email = ?", email).Error

	return user, err
}

func (r *authRepository) CreateUser(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) EmailExists(email string) bool {
	var count int64
	err := r.db.Model(&entities.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		log.Println(err)
	}

	return count > 0
}
