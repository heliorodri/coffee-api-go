package repository

import (
	models "coffee-api-go/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]*models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	CreateProduct(p *models.Product) error
	UpdateProduct(p *models.Product) error
	DeleteProduct(id uint) error
}

type productRepository struct {
	db gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: *db,
	}
}

func (repository *productRepository) GetAllProducts() ([]*models.Product, error) {
	var products []*models.Product
	err := repository.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (repository *productRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := repository.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repository *productRepository) CreateProduct(p *models.Product) error {
	return repository.db.Create(p).Error
}

func (repository *productRepository) UpdateProduct(p *models.Product) error {
	return repository.db.Save(p).Error
}

func (repository *productRepository) DeleteProduct(id uint) error {
	return repository.db.Delete(&models.Product{}, id).Error
}