package service

import (
	model "coffee-api-go/model/customer"
	repository "coffee-api-go/repository/customer"
	"coffee-api-go/utils"
	"errors"
)

type CustomerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{
		repo: *repo,
	}
}

func (s *CustomerService) GetAll() ([]*model.Customer, error) {
	customers, err := s.repo.GetAll()

	if err != nil {
		return nil, errors.New("FAILED to get all customers. Error: " + err.Error())
	}

	return customers, nil
}

func (s *CustomerService) GetByID(id uint) (*model.Customer, error) {
	customer, err := s.repo.GetByID(id)

	if err != nil {
		return nil, errors.New("FAILED to get customer with id " + utils.FormatId(id) + ".\n Error: " + err.Error())
	}

	return customer, nil
}

func (s *CustomerService) Create(p *model.Customer) (*model.Customer, error) {
	customer, err := s.repo.Create(p)

	if err != nil {
		return nil, errors.New("FAILED to create customer. Error: " + err.Error())
	}

	return customer, nil
}

func (s *CustomerService) Update(p *model.Customer) (*model.Customer, error) {
	customer, err := s.repo.Update(p)

	if err != nil {
		return nil, errors.New("FAILED to update customer. Error: " + err.Error())
	}

	return customer, nil
}
func (s *CustomerService) Delete(id uint) error {
	err := s.repo.Delete(id)

	if err != nil {
		return errors.New("FAILED to delete customer with id " + utils.FormatId(id) + ". Error: " + err.Error())
	}

	return nil
}
