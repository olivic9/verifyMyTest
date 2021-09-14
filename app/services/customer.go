package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"verifyMyTest/m/app/models"
	"verifyMyTest/m/app/repositories/contracts"
	"verifyMyTest/m/util"
)

type CustomerService struct {
	repo contracts.CustomerRepository
}

func NewCustomerService(r contracts.CustomerRepository) *CustomerService {
	return &CustomerService{
		repo: r,
	}
}

func (s *CustomerService) FindBy(findParameters util.FindParameters, parameters util.Parameters) *util.Data {
	return s.repo.FindBy(findParameters, parameters)
}

func (s *CustomerService) FindAll(parameters util.Parameters) *util.Data {
	return s.repo.FindAll(parameters)
}

func (s *CustomerService) Create(customer *models.Customer) (models.CustomerResponse, error) {
	hashedPassword, _ := HashPassword(customer.Password)
	customer.Password = hashedPassword
	r, err := s.repo.Create(customer)
	return r, err
}

func (s *CustomerService) Update(customer *models.Customer, id int) (error, int64) {
	return s.repo.Update(customer, id)
}

func (s *CustomerService) Delete(id int) (error, int64) {
	return s.repo.Delete(id)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}
