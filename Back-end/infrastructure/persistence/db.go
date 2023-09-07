package persistence

import (
	"database/sql"
	"fmt"
	"watchWebsite/domain/repository"

	_ "github.com/lib/pq"
)

// Database is a custom struct that holds the database connection pool
type Database struct {
	db *sql.DB
}

// Close closes the database connection
func (db *Database) Close() error {
	return db.db.Close()
}

func (db *Database) Begin() (*sql.Tx, error) {
	return db.db.Begin()
}

type Repositories struct {
	Watch       repository.WatchRepository
	Customer    repository.CustomerRepository
	Ward        repository.WardRepository
	District    repository.DistrictRepository
	Province    repository.ProvinceRepository
	Order       repository.OrderRepository
	OrderDetail repository.OrderDetailRepository
	Employee    repository.EmployeeRepository
	Brand       repository.BrandRepository
	TypeOfWatch repository.TypeOfWatchRepository
	DB          *Database
}

func NewRepositories(DbHost, DbPort, DbUser, DbName, DbPassword string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := sql.Open("postgres", DBURL)

	if err != nil {
		return nil, err
	}
	database := &Database{
		db: db,
	}

	return &Repositories{
		Watch:       NewWatchRepository(database),
		Customer:    NewCustomerRepository(database),
		Ward:        NewWardRepository(database),
		District:    NewDistrictRepository(database),
		Province:    NewProvinceRepository(database),
		Order:       NewOrderRepository(database),
		OrderDetail: NewOrderDetailRepository(database),
		Employee:    NewEmployeeRepository(database),
		Brand:       NewBrandRepository(database),
		TypeOfWatch: NewTypeOfWatchRepository(database),
		DB:          database,
	}, nil
}

// Close the  database connection
func (s *Repositories) Close() error {
	return s.DB.Close()
}
