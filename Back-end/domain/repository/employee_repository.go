package repository

import "watchWebsite/domain/entity"

type EmployeeRepository interface {
	GetAllEmployee() ([]*entity.Employee, error)
	GetEmployeeByIdAndPassword(id *uint32, password string) (*entity.Employee, error)
	CreateEmployee(*entity.Employee) (*uint32, error)
	UpdateEmployee(*entity.Employee) error
	DeleteEmployee(*uint32) error
}
