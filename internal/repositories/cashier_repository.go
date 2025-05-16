package repositories

import (
	"point-of-sales-app/internal/entities"

	"gorm.io/gorm"
)

type CashierRepository interface {
	CreateTransaction(transaction *entities.Transaction) error
	GetTransaction(id int) (*entities.Transaction, error)
	GetTransactions() ([]entities.Transaction, error)
	GetProductDetail(id int) (*entities.Product, error)
	UpdateStock(id int, updateData map[string]interface{}) error
}

type cashierRepository struct {
	db *gorm.DB
}

func NewCashierRepository(db *gorm.DB) CashierRepository {
	return &cashierRepository{
		db: db,
	}
}

func (r *cashierRepository) CreateTransaction(transaction *entities.Transaction) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entities.Product{}).Where("id = ?", transaction.ProductID).Update("quantity", gorm.Expr("quantity - ?", transaction.Quantity)).Error; err != nil {
			return err
		}

		if err := tx.Create(transaction).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *cashierRepository) GetTransaction(id int) (*entities.Transaction, error) {
	var transaction entities.Transaction
	err := r.db.First(&transaction, id).Error
	return &transaction, err
}

func (r *cashierRepository) GetTransactions() ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}

func (r *cashierRepository) GetProductDetail(id int) (*entities.Product, error) {
	var product entities.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *cashierRepository) UpdateStock(id int, updateData map[string]interface{}) error {
	return r.db.Model(&entities.Product{}).Where("id = ?", id).Updates(updateData).Error
}
