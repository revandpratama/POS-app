package services

import (
	"point-of-sales-app/internal/dto"
	"point-of-sales-app/internal/entities"
	"point-of-sales-app/internal/repositories"
)

type CashierService interface {
	CreateTransaction(req *dto.TransactionRequest) error
	UpdateStock(id int, newNumber int) error
	GetTransactionList() (*[]dto.TransactionResponse, error)
}

type cashierService struct {
	repo repositories.CashierRepository
}

func NewCashierService(repo repositories.CashierRepository) CashierService {
	return &cashierService{repo: repo}
}

func (s *cashierService) CreateTransaction(req *dto.TransactionRequest) error {

	product, err := s.repo.GetProductDetail(req.ProductID)
	if err != nil {
		return err
	}

	transaction := &entities.Transaction{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Total:     float64(req.Quantity) * float64(product.Price),
	}
	return s.repo.CreateTransaction(transaction)
}

func (s *cashierService) UpdateStock(id int, newNumber int) error {

	updateData := map[string]interface{}{
		"quantity": newNumber,
	}
	return s.repo.UpdateStock(id, updateData)
}

func (s *cashierService) GetTransactionList() (*[]dto.TransactionResponse, error) {
	transactions, err := s.repo.GetTransactions()
	if err != nil {
		return nil, err
	}

	var transactionsResponse = make([]dto.TransactionResponse, len(transactions))

	for i, transaction := range transactions {
		transactionsResponse[i] = dto.TransactionResponse{
			ID:        transaction.ID,
			ProductID: transaction.ProductID,
			Quantity:  transaction.Quantity,
			Total:     transaction.Total,
			CreatedAt: transaction.CreatedAt,
		}
	}

	return &transactionsResponse, nil
}
