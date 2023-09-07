package entity

type Watch struct {
	WatchId       *uint32 `json:"watchId"`
	WatchName     string  `json:"watchName"`
	Price         *uint32 `json:"price"`
	Image         string  `json:"image"`
	Status        *int32  `json:"status"`
	Quantity      *uint32 `json:"quantity"`
	BrandId       *uint32 `json:"brandId"`
	TypeOfWatchId *uint32 `json:"typeOfWatchId"`
}
