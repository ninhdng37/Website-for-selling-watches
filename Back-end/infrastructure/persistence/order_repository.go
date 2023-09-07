package persistence

import (
	"database/sql"
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type OrderRepo struct {
	db *Database
}

func NewOrderRepository(db *Database) *OrderRepo {
	return &OrderRepo{db}
}

var _ repository.OrderRepository = &OrderRepo{}

func (o *OrderRepo) CreateOrder(tx *sql.Tx, order *entity.Order) (*uint32, error) {
	// tx, err := o.db.Begin()
	// if err != nil {
	// 	return nil, err
	// }
	// defer tx.Rollback()
	query := `	
	INSERT INTO orders (
	customer_id,
	province_id,
	district_id,
	ward_id,
	apartment_number)
	VALUES ($1, $2, $3, $4, $5) RETURNING order_id`
	var insertedID uint32
	err := tx.QueryRow(
		query,
		order.CustomerId,
		order.ProvinceId,
		order.DistrictId,
		order.WardId,
		order.ApartmentNumber).Scan(&insertedID)

	if err != nil {
		return nil, err
	}
	// tx.Commit()
	return &insertedID, nil
}

func (o *OrderRepo) GetOrderByCustomerID(customerID *uint32) ([]*entity.Order, error) {
	rows, err := o.db.db.Query(`	SELECT *
									FROM orders 
									WHERE customer_id = $1`, *customerID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []*entity.Order

	for rows.Next() {
		order := &entity.Order{}
		if err := rows.Scan(
			&order.OrderId,
			&order.OrderDate,
			&order.CustomerId,
			&order.EmployeeId,
			&order.InvoiceId,
			&order.Status,
			&order.ProvinceId,
			&order.DistrictId,
			&order.WardId,
			&order.ApartmentNumber); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *OrderRepo) GetAllOrders() ([]*entity.OrderResponse, error) {
	rows, err := o.db.db.Query(`	SELECT O.ORDER_ID, O.ORDER_DATE, O.CUSTOMER_NAME, E.fullname,
	O.INVOICE_ID, O.STATUS, O.PROVINCE_NAME, O.DISTRICT_NAME, O.WARD_NAME, O.APARTMENT_NUMBER
	FROM
	(SELECT * FROM
	(SELECT * FROM
	(SELECT * FROM 
	(SELECT order_id, order_date, customer_name, EMPLOYEE_ID, invoice_id, status, o.province_id, o.district_id,
	 o.ward_id, o.apartment_number
	FROM orders o INNER JOIN customers cus ON o.customer_id = cus.customer_id) O 
	INNER JOIN provinces pr ON pr.province_ID = O.PROVINCE_ID) O
	INNER JOIN DISTRICTS D ON D.DISTRICT_ID = O.DISTRICT_ID) O 
	INNER JOIN WARDS W ON W.WARD_ID = O.WARD_ID) O 
	LEFT JOIN EMPLOYEES E ON E.EMPLOYEE_ID = O.EMPLOYEE_ID
	ORDER BY order_id `)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []*entity.OrderResponse

	for rows.Next() {
		order := &entity.OrderResponse{}
		if err := rows.Scan(
			&order.OrderID,
			&order.OrderDate,
			&order.CustomerName,
			&order.EmployeeName,
			&order.InvoiceID,
			&order.Status,
			&order.ProviceName,
			&order.DistrictName,
			&order.WardName,
			&order.ApartmentNumber); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *OrderRepo) UpdateOrderByCustomer(orderID *uint32) error {
	tx, err := o.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `	
	UPDATE orders
	SET 
	status = 0
	WHERE order_id = $1`
	_, err = tx.Exec(
		query,
		*orderID)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (o *OrderRepo) UpdateOrderStatus(orderID, employeeID, status *uint32) error {

	tx, err := o.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `	
	UPDATE orders
	SET 
	status = $1,
	employee_id = $2
	WHERE order_id = $3`
	_, err = tx.Exec(
		query,
		status,
		employeeID,
		orderID)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
