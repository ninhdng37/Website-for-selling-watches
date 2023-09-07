package repository

import "watchWebsite/domain/entity"

type CustomerRepository interface {
	GetCustomerById(customerId uint32) (*entity.Customer, error)
	GetCustomerByEmail(email string) (*entity.Customer, error)
	GetCustomerByPhoneNumber(phoneNumber string) (*entity.Customer, error)
	GetCustomerByEmailAndPassword(email, password string) (*entity.Customer, error)
	CreateCustomer(customer *entity.Customer) error
	UpdateCustomer(customer *entity.Customer) error
}
