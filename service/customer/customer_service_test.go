package service

import (
	"reflect"
	"testing"

	models "coffee-api-go/model/customer"
)

func TestGetAll(t *testing.T) {
	expected := []*models.Customer{{Name: "John Doe"}, {Name: "Jane Smith"}}
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

	if !reflect.DeepEqual(customers, expected) {
		t.Errorf("Expected %v but got %v", expected, customers)
	}
}

// Define a mock repository struct that implements the repository interface
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
