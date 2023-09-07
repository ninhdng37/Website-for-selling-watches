package application

import (
	"database/sql"
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type orderApp struct {
	o repository.OrderRepository
}

var _ OrderAppInterface = &orderApp{}

type OrderAppInterface interface {
	CreateOrder(tx *sql.Tx, order *entity.Order) (*uint32, error)
	GetOrderByCustomerID(customerID *uint32) ([]*entity.Order, error)
	GetAllOrders() ([]*entity.OrderResponse, error)
	UpdateOrderByCustomer(orderID *uint32) error
	UpdateOrderStatus(orderID, employeeID, status *uint32) error
}

func (o *orderApp) CreateOrder(tx *sql.Tx, order *entity.Order) (*uint32, error) {
	return o.o.CreateOrder(tx, order)
}

func (o *orderApp) GetOrderByCustomerID(customerID *uint32) ([]*entity.Order, error) {
	return o.o.GetOrderByCustomerID(customerID)
}

func (o *orderApp) GetAllOrders() ([]*entity.OrderResponse, error) {
	return o.o.GetAllOrders()
}

func (o *orderApp) UpdateOrderByCustomer(orderID *uint32) error {
	return o.o.UpdateOrderByCustomer(orderID)
}

func (o *orderApp) UpdateOrderStatus(orderID, employeeID, status *uint32) error {
	return o.o.UpdateOrderStatus(orderID, employeeID, status)
}
