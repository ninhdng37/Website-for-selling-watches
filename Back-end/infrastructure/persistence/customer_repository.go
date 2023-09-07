package persistence

import (
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
	"watchWebsite/infrastructure/security"
)

type CustomerRepo struct {
	db *Database
}

func NewCustomerRepository(db *Database) *CustomerRepo {
	return &CustomerRepo{db}
}

var _ repository.CustomerRepository = &CustomerRepo{}

func (r *CustomerRepo) GetCustomerById(customerId uint32) (*entity.Customer, error) {
	//Perform the SELECT query
	rows := r.db.db.QueryRow(`SELECT *
								FROM customers
								WHERE customer_id = $1`, customerId)

	customer := &entity.Customer{}
	err := rows.Scan(
		customer.CustomerId,
		customer.CustomerName,
		customer.Email,
		customer.PhoneNumber,
		customer.ApartmentNumber,
		customer.Password,
		customer.ProvinceId,
		customer.DistrictId,
		customer.WardId,
		customer.IsVerified,
	)

	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *CustomerRepo) GetCustomerByEmail(email string) (*entity.Customer, error) {
	//Perform the SELECT query
	rows := r.db.db.QueryRow(`SELECT *
								FROM customers
								WHERE email = $1`, email)

	customer := &entity.Customer{}
	err := rows.Scan(
		&customer.CustomerId,
		&customer.CustomerName,
		&customer.Email,
		&customer.PhoneNumber,
		&customer.ApartmentNumber,
		&customer.Password,
		&customer.ProvinceId,
		&customer.DistrictId,
		&customer.WardId,
		&customer.IsVerified)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepo) GetCustomerByPhoneNumber(phoneNumber string) (*entity.Customer, error) {
	//Perform the SELECT query
	rows := r.db.db.QueryRow(`SELECT *
								FROM customers
								WHERE phone_number = $1`, phoneNumber)

	customer := &entity.Customer{}
	err := rows.Scan(
		&customer.CustomerId,
		&customer.CustomerName,
		&customer.Email,
		&customer.PhoneNumber,
		&customer.ApartmentNumber,
		&customer.Password,
		&customer.ProvinceId,
		&customer.DistrictId,
		&customer.WardId,
		&customer.IsVerified)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepo) GetCustomerByEmailAndPassword(email, password string) (*entity.Customer, error) {
	password = security.Hash(password)
	//Perform the SELECT query
	rows := r.db.db.QueryRow(`SELECT *
								FROM customers
								WHERE email = $1 AND password = $2`, email, password)

	customer := &entity.Customer{}
	err := rows.Scan(
		&customer.CustomerId,
		&customer.CustomerName,
		&customer.Email,
		&customer.PhoneNumber,
		&customer.ApartmentNumber,
		&customer.Password,
		&customer.ProvinceId,
		&customer.DistrictId,
		&customer.WardId,
		&customer.IsVerified)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *CustomerRepo) CreateCustomer(customer *entity.Customer) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `	
	INSERT INTO customers (
	customer_name,
	email,
	phone_number,
	apartment_number,
	password,
	province_id,
	district_id,
	ward_id,
	is_verified)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	customer.Password = security.Hash(customer.Password)

	if err != nil {
		return err
	}
	_, err = tx.Exec(
		query,
		customer.CustomerName,
		customer.Email,
		customer.PhoneNumber,
		customer.ApartmentNumber,
		customer.Password,
		customer.ProvinceId,
		customer.DistrictId,
		customer.WardId,
		customer.IsVerified)

	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (r *CustomerRepo) UpdateCustomer(customer *entity.Customer) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `	
	UPDATE customers
	SET 
	customer_name = CASE WHEN $1 = '-1' THEN customer_name ELSE $1 END,
	email = CASE WHEN $2 = '-1' THEN email ELSE $2 END,
	phone_number = CASE WHEN $3 = '-1' THEN phone_number ELSE $3 END,
	apartment_number = CASE WHEN $4 = '-1' THEN apartment_number ELSE $4 END,
	password = CASE WHEN $5 = '-1' THEN password ELSE $5 END,
	province_id = CASE WHEN $6 = 0 THEN province_id ELSE $6 END,
	district_id = CASE WHEN $7 = 0 THEN district_id ELSE $7 END,
	ward_id = CASE WHEN $8 = 0 THEN ward_id ELSE $8 END,
	is_verified = $9
	WHERE customer_id = $10`
	_, err = tx.Exec(
		query,
		customer.CustomerName,
		customer.Email,
		customer.PhoneNumber,
		customer.ApartmentNumber,
		customer.Password,
		customer.ProvinceId,
		customer.DistrictId,
		customer.WardId,
		customer.IsVerified,
		customer.CustomerId)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
