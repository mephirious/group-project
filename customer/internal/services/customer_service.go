package services

import (
	"strconv"

	"github.com/mephirious/group-project/services/customer/internal/models"
	"github.com/mephirious/group-project/services/customer/internal/repository"
)

type CustomerService struct {
	Repo *repository.CustomerRepository
}

// GetCustomerProfile получает данные профиля по ID
func (s *CustomerService) GetCustomerProfile(customerID string) (*models.Customer, error) {
	customerIDint, err := strconv.Atoi(customerID)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetCustomerByID(customerIDint)
}
