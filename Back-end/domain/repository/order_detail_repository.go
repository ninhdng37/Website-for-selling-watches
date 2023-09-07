package repository

import (
	"database/sql"
	"watchWebsite/domain/entity"
)

type OrderDetailRepository interface {
	CreateOrderDetail(tx *sql.Tx, orderDetail *entity.OrderDetail) error
	GetOrderDetailByCustomerEmail(email string) ([]*entity.OrderDetailResponse, error)
	GetOrderDetailByOrderID(orderID *uint32) ([]*entity.OrderDetailResponse, error)
}
