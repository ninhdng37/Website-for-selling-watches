package persistence

import (
	"errors"
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type EmployeeRepo struct {
	db *Database
}

func NewEmployeeRepository(db *Database) *EmployeeRepo {
	return &EmployeeRepo{db}
}

var _ repository.EmployeeRepository = &EmployeeRepo{}

func (e *EmployeeRepo) GetAllEmployee() ([]*entity.Employee, error) {
	rows, err := e.db.db.Query(`	SELECT employee_id, 
									fullname, 
									identity_number, 
									position, 
									email,
									phone_number
									FROM employees
									ORDER BY employee_id`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees []*entity.Employee

	for rows.Next() {
		employee := &entity.Employee{}
		if err := rows.Scan(
			&employee.EmployeeID,
			&employee.Fullname,
			&employee.IdentityNumber,
			&employee.Position,
			&employee.Email,
			&employee.PhoneNumber); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (e *EmployeeRepo) GetEmployeeByIdAndPassword(id *uint32, password string) (*entity.Employee, error) {
	rows := e.db.db.QueryRow(`SELECT *
								FROM employees
								WHERE employee_id = $1 AND password = $2`, id, password)

	employee := &entity.Employee{}
	err := rows.Scan(
		&employee.EmployeeID,
		&employee.Fullname,
		&employee.IdentityNumber,
		&employee.Position,
		&employee.Email,
		&employee.PhoneNumber,
		&employee.Password,
	)

	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employee, nil
}

func (e *EmployeeRepo) CreateEmployee(employee *entity.Employee) (*uint32, error) {
	tx, err := e.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	query := `	
	INSERT INTO employees (
	fullname,
	identity_number,
	position,
	email,
	phone_number,
	password)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING employee_id`
	var employeeId uint32
	err = tx.QueryRow(
		query,
		employee.Fullname,
		employee.IdentityNumber,
		employee.Position,
		employee.Email,
		employee.PhoneNumber,
		employee.Password).Scan(&employeeId)

	if err != nil {
		return nil, err
	}
	tx.Commit()
	return &employeeId, nil
}

func (e *EmployeeRepo) UpdateEmployee(employee *entity.Employee) error {
	tx, err := e.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	employeeExisted := e.isExist(employee)

	if employeeExisted {
		var ErrEmployeeExists = errors.New("other employee had your information typed by you")
		return ErrEmployeeExists
	}

	query := `	
	UPDATE employees
	SET  
	fullname = CASE WHEN $1 = '-1' THEN fullname ELSE $1 END,
	identity_number = CASE WHEN $2 = '-1' THEN identity_number ELSE $2 END,
	position = CASE WHEN $3 = '-1' THEN position ELSE $3 END,
	email = CASE WHEN $4 = '-1' THEN email ELSE $4 END,
	phone_number = CASE WHEN $5 = '-1' THEN phone_number ELSE $5 END
	WHERE employee_id = $6`

	_, err = tx.Exec(
		query,
		employee.Fullname,
		employee.IdentityNumber,
		employee.Position,
		employee.Email,
		employee.PhoneNumber,
		employee.EmployeeID)

	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (e *EmployeeRepo) isExist(employee *entity.Employee) bool {
	// Check uniqueness for IdentityNumber, Email, and PhoneNumber
	query := `
		SELECT COUNT(*) FROM employees
		WHERE (identity_number = $1 OR email = $2 OR phone_number = $3) AND employee_id != $4
	`
	var count int
	err := e.db.db.QueryRow(query, employee.IdentityNumber, employee.Email,
		employee.PhoneNumber, employee.EmployeeID).Scan(&count)
	if err != nil {
		// fmt.Println("Database error:", err)
		return false
	}

	return count != 0
}

func (e *EmployeeRepo) DeleteEmployee(employeeID *uint32) error {
	tx, err := e.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var count int
	query := "SELECT COUNT(*) FROM orders WHERE employee_id = $1"
	err = tx.QueryRow(query, employeeID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("can not delete employee because this employee approved orders")
	}
	_, err = e.db.db.Exec("DELETE FROM employees WHERE employee_id = $1", employeeID)

	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
