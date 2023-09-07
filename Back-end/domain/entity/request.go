package entity

type OrderRequest struct {
	ProvinceId      *uint32    `json:"provinceId"`
	DistrictId      *uint32    `json:"districtId"`
	WardId          *uint32    `json:"wardId"`
	ApartmentNumber string     `json:"apartmentNumber"`
	Items           []CartData `json:"items"`
}

type CartData struct {
	Title     string  `json:"title"`
	UnitPrice *uint32 `json:"unitPrice"`
	Quantity  *uint32 `json:"quantity"`
}

type EmployeeRequest struct {
	Fullname       string `json:"fullname"`
	IdentityNumber string `json:"identityNumber"`
	Position       string `json:"position"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	Password       string `json:"password"`
}
