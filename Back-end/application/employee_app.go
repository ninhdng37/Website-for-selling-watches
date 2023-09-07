package application

import (
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type employeeApp struct {
	er repository.EmployeeRepository
}

var _ EmployeeAppInterface = &employeeApp{}

type EmployeeAppInterface interface {
	CreateEmployee(*entity.Employee) (*uint32, error)
	GetEmployeeByIdAndPassword(id *uint32, password string) (*entity.Employee, error)
	GetAllEmployee() ([]*entity.Employee, error)
	UpdateEmployee(employee *entity.Employee) error
	DeleteEmployee(employeeID *uint32) error
}

func (e *employeeApp) GetEmployeeByIdAndPassword(id *uint32, password string) (*entity.Employee, error) {
	return e.er.GetEmployeeByIdAndPassword(id, password)
}

func (e *employeeApp) CreateEmployee(employee *entity.Employee) (*uint32, error) {
	return e.er.CreateEmployee(employee)
}

func (e *employeeApp) GetAllEmployee() ([]*entity.Employee, error) {
	return e.er.GetAllEmployee()
}

func (e *employeeApp) UpdateEmployee(employee *entity.Employee) error {
	return e.er.UpdateEmployee(employee)
}

func (e employeeApp) DeleteEmployee(employeeID *uint32) error {
	return e.er.DeleteEmployee(employeeID)
}
