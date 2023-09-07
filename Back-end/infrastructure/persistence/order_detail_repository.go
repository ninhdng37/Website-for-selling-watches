package persistence

import (
	"database/sql"
	"fmt"
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type OrderDetailRepo struct {
	db *Database
}

func NewOrderDetailRepository(db *Database) *OrderDetailRepo {
	return &OrderDetailRepo{db}
}

var _ repository.OrderDetailRepository = &OrderDetailRepo{}

func (od *OrderDetailRepo) CreateOrderDetail(tx *sql.Tx, orderDetail *entity.OrderDetail) error {
	// tx, err := od.db.Begin()
	// if err != nil {
	// 	return err
	// }
	// defer tx.Rollback()
	query := `	
	INSERT INTO order_details (
	watch_id,
	order_id,
	quantity,
	unit_price)
	VALUES ($1, $2, $3, $4)`

	_, err := tx.Exec(
		query,
		orderDetail.WatchId,
		orderDetail.OrderId,
		orderDetail.Quantity,
		orderDetail.UnitPrice)

	if err != nil {
		fmt.Println(err)
		return err
	}

	// tx.Commit()
	return nil
}

func (od *OrderDetailRepo) GetOrderDetailByCustomerEmail(email string) ([]*entity.OrderDetailResponse, error) {
	rows, err := od.db.db.Query(`	SELECT w.watch_name, w.image,order_details.order_date, order_details.order_id, order_details.status,
	order_details.quantity, order_details.unit_price
	FROM watches w
	INNER JOIN (
		SELECT od.order_id, customer_orders.status,customer_orders.order_date
		, od.watch_id, od.quantity, od.unit_price
		FROM order_details od
		INNER JOIN (
			SELECT o.order_id, o.status, o.order_date
			FROM orders o
			WHERE o.customer_id = (
				SELECT c.customer_id
				FROM customers c
				WHERE c.email = $1
			)
		) AS customer_orders
		ON od.order_id = customer_orders.order_id
	) AS order_details
	ON w.watch_id = order_details.watch_id order by order_id`, email)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orderDetails []*entity.OrderDetailResponse

	for rows.Next() {
		orderDetail := &entity.OrderDetailResponse{}
		if err := rows.Scan(
			&orderDetail.WatchName,
			&orderDetail.WatchImage,
			&orderDetail.OrderDate,
			&orderDetail.OrderId,
			&orderDetail.Status,
			&orderDetail.Quantity,
			&orderDetail.UnitPrice); err != nil {
			return nil, err
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orderDetails, nil
}

func (od *OrderDetailRepo) GetOrderDetailByOrderID(orderID *uint32) ([]*entity.OrderDetailResponse, error) {

	rows, err := od.db.db.Query(`SELECT W.WATCH_NAME, OD.QUANTITY, OD.UNIT_PRICE 
	FROM ORDER_DETAILS OD
	INNER JOIN WATCHES W ON W.WATCH_ID = OD.WATCH_ID AND OD.ORDER_ID = $1
	ORDER BY w.watch_name`, orderID)

	if err != nil {
		return nil, err
	}

	var orders []*entity.OrderDetailResponse

	for rows.Next() {
		order := &entity.OrderDetailResponse{}
		if err := rows.Scan(
			&order.WatchName,
			&order.Quantity,
			&order.UnitPrice); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
