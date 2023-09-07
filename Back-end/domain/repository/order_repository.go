package repository

import (
	"database/sql"
	"watchWebsite/domain/entity"
)

type OrderRepository interface {
	CreateOrder(tx *sql.Tx, order *entity.Order) (*uint32, error)
	GetOrderByCustomerID(customerID *uint32) ([]*entity.Order, error)
	GetAllOrders() ([]*entity.OrderResponse, error)
	UpdateOrderByCustomer(orderID *uint32) error
	UpdateOrderStatus(orderID, employeeID, status *uint32) error
}
