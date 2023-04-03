package service

import (
	"errors"
	"reflect"
	"testing"

	models "coffee-api-go/model/customer"

	"github.com/stretchr/testify/assert"
)

func TestGetAllWithData(t *testing.T) {
	expected := []*models.Customer{
		{
			Name:  "John Doe",
			Email: "johndoe@mail.com",
		},
		{
			Name:  "Jane Smith",
			Email: "janesmith@mail.com",
		},
	}

	repo := &mockRepository{
		mockGetAll: func() ([]*models.Customer, error) {
			return expected, nil
		},
	}

	service := &CustomerService{repo: repo}
	customers, err := service.GetAll()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.True(t, reflect.DeepEqual(customers, expected))
}

func TestGetAllWithEmptyReturn(t *testing.T) {
	expected := []*models.Customer{}

	repo := &mockRepository{
		mockGetAll: func() ([]*models.Customer, error) {
			return expected, nil
		},
	}

	service := &CustomerService{repo: repo}
	customers, err := service.GetAll()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.True(t, reflect.DeepEqual(customers, expected))
}

func TestGetAllWithError(t *testing.T) {
	repo := &mockRepository{
		mockGetAll: func() ([]*models.Customer, error) {
			return nil, errors.New("Error connecting to database")
		},
	}

	service := &CustomerService{repo: repo}
	_, err := service.GetAll()

	assert.EqualError(t, err, "FAILED to get all customers. Error: Error connecting to database")
}

func TestGetByIDWithData(t *testing.T) {
	expected := &models.Customer{
		Name:  "John Doe",
		Email: "johndoe_get_by_id@mail.com",
	}

	repo := &mockRepository{
		mockGetByID: func() (*models.Customer, error) {
			return expected, nil
		},
	}

	service := &CustomerService{repo: repo}
	customer, err := service.GetByID(1)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.True(t, reflect.DeepEqual(customer, expected))
}

func TestGetByIdNoDataFound(t *testing.T) {
	repo := &mockRepository{
		mockGetByID: func() (*models.Customer, error) {
			return nil, errors.New("Customer not found - test")
		},
	}

	service := &CustomerService{repo: repo}
	_, err := service.GetByID(1)

	assert.EqualError(t, err, "FAILED to get customer with id 1. Error: Customer not found - test")
}

type mockRepository struct {
	mockGetAll  func() ([]*models.Customer, error)
	mockGetByID func() (*models.Customer, error)
	mockCreate  func() (*models.Customer, error)
	mockUpdate  func() (*models.Customer, error)
	mockDelete  func() error
}

func (repository *mockRepository) GetAll() ([]*models.Customer, error) {
	return repository.mockGetAll()
}

func (repository *mockRepository) GetByID(id uint) (*models.Customer, error) {
	return repository.mockGetByID()
}

func (repository *mockRepository) Create(p *models.Customer) (*models.Customer, error) {
	return repository.mockCreate()
}

func (repository *mockRepository) Update(p *models.Customer) (*models.Customer, error) {
	return repository.mockUpdate()
}

func (repository *mockRepository) Delete(id uint) error {
	return repository.mockDelete()
}
