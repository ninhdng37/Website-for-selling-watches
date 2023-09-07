package application

import (
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type customerApp struct {
	cr repository.CustomerRepository
}

var _ CustomerAppInterface = &customerApp{}

type CustomerAppInterface interface {
	GetCustomerById(id uint32) (*entity.Customer, error)
	GetCustomerByEmail(email string) (*entity.Customer, error)
	GetCustomerByPhoneNumber(phoneNumber string) (*entity.Customer, error)
	GetCustomerByEmailAndPassword(email, password string) (*entity.Customer, error)
	CreateCustomer(customer *entity.Customer) error
	UpdateCustomer(customer *entity.Customer) error
}

func (c *customerApp) GetCustomerById(customerId uint32) (*entity.Customer, error) {
	return c.cr.GetCustomerById(customerId)
}

func (c *customerApp) GetCustomerByEmail(customerEmail string) (*entity.Customer, error) {
	return c.cr.GetCustomerByEmail(customerEmail)
}

func (c *customerApp) GetCustomerByPhoneNumber(phoneNumber string) (*entity.Customer, error) {
	return c.cr.GetCustomerByPhoneNumber(phoneNumber)
}

func (c *customerApp) GetCustomerByEmailAndPassword(email, password string) (*entity.Customer, error) {
	return c.cr.GetCustomerByEmailAndPassword(email, password)
}

func (c *customerApp) CreateCustomer(customer *entity.Customer) error {
	return c.cr.CreateCustomer(customer)
}

func (c *customerApp) UpdateCustomer(customer *entity.Customer) error {
	return c.cr.UpdateCustomer(customer)
}
