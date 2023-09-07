package entity

type OrderDetail struct {
	WatchId   *uint32 `json:"watchId"`
	OrderId   *uint32 `json:"orderId"`
	Quantity  *uint32 `json:"quantity"`
	UnitPrice *uint32 `json:"unitPrice"`
}
