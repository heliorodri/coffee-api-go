package repository

import (
	models "coffee-api-go/model/customer"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetAll() ([]*models.Customer, error)
	GetByID(id uint) (*models.Customer, error)
	Create(p *models.Customer) (*models.Customer, error)
	Update(p *models.Customer) (*models.Customer, error)
	Delete(id uint) error
}

type customerRepository struct {
	db gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: *db,
	}
}

func (repository *customerRepository) GetAll() ([]*models.Customer, error) {
	var customers []*models.Customer

	err := repository.db.Find(&customers).Error

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (repository *customerRepository) GetByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	err := repository.db.First(&customer, id).Error

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (repository *customerRepository) Create(p *models.Customer) (*models.Customer, error) {
	repository.db.Save(p)

	var savedCustomer models.Customer
	err := repository.db.First(&savedCustomer, p.ID).Error

	if err != nil {
		return nil, err
	}

	return &savedCustomer, nil
}

func (repository *customerRepository) Update(p *models.Customer) (*models.Customer, error) {
	repository.db.Save(p)

	var updatedCustomer models.Customer
	err := repository.db.First(&updatedCustomer, p.ID).Error

	if err != nil {
		return nil, err
	}

	return &updatedCustomer, nil
}

func (repository *customerRepository) Delete(id uint) error {
	var customer models.Customer

	err := repository.db.Delete(&customer, id).Error

	if err != nil {
		return err
	}

	return nil
}
