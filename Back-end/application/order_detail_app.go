package application

import (
	"database/sql"
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type orderDetailApp struct {
	od repository.OrderDetailRepository
}

var _ OrderDetailAppInterface = &orderDetailApp{}

type OrderDetailAppInterface interface {
	CreateOrderDetail(tx *sql.Tx, orderDetail *entity.OrderDetail) error
	GetOrderDetailByCustomerEmail(email string) ([]*entity.OrderDetailResponse, error)
	GetOrderDetailByOrderID(orderID *uint32) ([]*entity.OrderDetailResponse, error)
}

func (od *orderDetailApp) CreateOrderDetail(tx *sql.Tx, orderDetail *entity.OrderDetail) error {
	return od.od.CreateOrderDetail(tx, orderDetail)
}

func (od *orderDetailApp) GetOrderDetailByCustomerEmail(email string) ([]*entity.OrderDetailResponse, error) {
	return od.od.GetOrderDetailByCustomerEmail(email)
}

func (od *orderDetailApp) GetOrderDetailByOrderID(orderID *uint32) ([]*entity.OrderDetailResponse, error) {
	return od.od.GetOrderDetailByOrderID(orderID)
}
