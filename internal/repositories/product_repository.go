package repositories

import (
	"point-of-sales-app/internal/entities"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts() ([]entities.Product, error)
	GetProduct(id int) (*entities.Product, error)
	CreateProduct(product *entities.Product) error
	UpdateProduct(id int, updateData map[string]interface{}) error
	DeleteProduct(id int) error
}

type productRepository struct {
	db *gorm.DB
}


func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}


func (r *productRepository) GetProducts() ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Find(&products).Error

	return products, err
}


func (r *productRepository) GetProduct(id int) (*entities.Product, error) {
	var product entities.Product
	err := r.db.First(&product, id).Error
	return &product, err
}


func (r *productRepository) CreateProduct(product *entities.Product) error {
	return r.db.Create(product).Error
}


func (r *productRepository) UpdateProduct(id int, updateData map[string]interface{}) error {
	return r.db.Model(&entities.Product{}).Where("id = ?", id).Updates(updateData).Error
}


func (r *productRepository) DeleteProduct(id int) error {
	return r.db.Delete(&entities.Product{}, id).Error
}

