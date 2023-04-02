package service

import (
	models "coffee-api-go/model"
	"coffee-api-go/repository"
	"errors"
	"strconv"
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
	products, err := s.repo.GetAllProducts()

	if err != nil {
		return nil, errors.New("FAILED to get all products.\n Error: " + err.Error())
	}

	return products, nil
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	product, err := s.repo.GetProductByID(id)

	if err != nil {
		return nil, errors.New("FAILED to get product with id " + strconv.FormatUint(uint64(id), 10) + ".\n Error: " + err.Error())
	}

	return product, nil
}

func (s *ProductService) CreateProduct(p *models.Product) (*models.Product, error) {
	product, err := s.repo.CreateProduct(p)

	if err != nil {
		return nil, errors.New("FAILED to create product.\n Error: " + err.Error())
	}

	return product, nil
}

func (s *ProductService) UpdateProduct(p *models.Product) (*models.Product, error) {
	product, err := s.repo.UpdateProduct(p)

	if err != nil {
		return nil, errors.New("FAILED to update product.\n Error: " + err.Error())
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	if err := s.repo.DeleteProduct(id); err != nil {
		return errors.New("FAILED to delete product with id " + strconv.FormatUint(uint64(id), 10) + ".\n Error: " + err.Error())
	}

	return nil
}
