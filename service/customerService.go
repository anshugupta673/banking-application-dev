package service

import(
	"github.com/anshugupta673/go-banking-application/domain"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain,CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error){
	return s.repo.FindAll()
}

func NewCustomerService(retpository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

