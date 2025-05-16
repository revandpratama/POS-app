package services

import (
	"point-of-sales-app/internal/dto"
	"point-of-sales-app/internal/entities"
	"point-of-sales-app/internal/repositories"
)

type ProductService interface {
	GetProducts() ([]dto.ProductResponse, error)
	GetProduct(id int) (*dto.ProductResponse, error)
	CreateProduct(product *dto.ProductRequest) error
	UpdateProduct(id int, product *dto.ProductRequest) error
	DeleteProduct(id int) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (p *productService) GetProducts() ([]dto.ProductResponse, error) {
	products, err := p.repo.GetProducts()
	if err != nil {
		return nil, err
	}

	var productsResponse = make([]dto.ProductResponse, len(products))

	for i, product := range products {
		productsResponse[i] = dto.ProductResponse{
			ID:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		}
	}

	return productsResponse, nil
}

func (p *productService) GetProduct(id int) (*dto.ProductResponse, error) {
	product, err := p.repo.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}, nil
}

func (p *productService) CreateProduct(product *dto.ProductRequest) error {
	return p.repo.CreateProduct(&entities.Product{
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	})
}

func (p *productService) UpdateProduct(id int, product *dto.ProductRequest) error {

	updateProduct := make(map[string]any)

	if product.Name != "" {
		updateProduct["name"] = product.Name
	}
	if product.Price != 0 {
		updateProduct["price"] = product.Price
	}
	if product.Quantity != 0 {
		updateProduct["quantity"] = product.Quantity
	}

	return p.repo.UpdateProduct(id, updateProduct)
}


func (p *productService) DeleteProduct(id int) error {
	return p.repo.DeleteProduct(id)
}

