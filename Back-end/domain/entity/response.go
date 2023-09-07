package entity

import (
	"strings"
	"time"
)

type OrderDetailResponse struct {
	WatchName  string
	WatchImage string
	OrderDate  *time.Time
	OrderId    *uint32
	Status     *uint32
	Quantity   *uint32
	UnitPrice  *uint32
}

type OrderResponse struct {
	OrderID         *uint32    `json:"orderID"`
	OrderDate       *time.Time `json:"orderDate"`
	OrderDateString string     `json:"orderDateString"`
	CustomerName    string     `json:"customerName"`
	EmployeeName    *string    `json:"employeeName"`
	InvoiceID       *uint32    `json:"invoiceID"`
	Status          *uint32    `json:"status"`
	ProviceName     string     `json:"provinceName"`
	DistrictName    string     `json:"districtName"`
	WardName        string     `json:"wardName"`
	ApartmentNumber string     `json:"apartmentNumber"`
}

func (o *OrderResponse) TrimSpaces() {
	o.CustomerName = strings.TrimSpace(o.CustomerName)
	if o.EmployeeName != nil {
		*o.EmployeeName = strings.TrimSpace(*o.EmployeeName)
	}
	o.OrderDateString = o.OrderDate.Format("02/01/2006 15:04:05")
	o.ProviceName = strings.TrimSpace(o.ProviceName)
	o.DistrictName = strings.TrimSpace(o.DistrictName)
	o.WardName = strings.TrimSpace(o.WardName)
	o.ApartmentNumber = strings.TrimSpace(o.ApartmentNumber)
}
