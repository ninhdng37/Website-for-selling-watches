package entity

import "time"

type Order struct {
	OrderId         *uint32    `json:"orderId"`
	OrderDate       *time.Time `json:"orderDate"`
	CustomerId      *uint32    `json:"customerId"`
	EmployeeId      *uint32    `json:"employeeId"`
	InvoiceId       *uint32    `json:"invoiceId"`
	Status          *uint32    `json:"status"`
	ProvinceId      *uint32    `json:"provinceId"`
	DistrictId      *uint32    `json:"districtId"`
	WardId          *uint32    `json:"wardId"`
	ApartmentNumber string     `json:"ApartmentNumber"`
}
