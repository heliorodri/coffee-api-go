package service

import (
	models "coffee-api-go/model"
	"coffee-api-go/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: *repo,
	}
}

func (s *ProductService) GetAllProducts() ([]*models.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) CreateProduct(p *models.Product) error {
	return s.repo.CreateProduct(p)
}

func (s *ProductService) UpdateProduct(p *models.Product) error {
	return s.repo.UpdateProduct(p)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}
