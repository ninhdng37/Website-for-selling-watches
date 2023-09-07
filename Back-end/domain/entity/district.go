package entity

type District struct {
	DistrictId   *uint32 `json:"districtId"`
	DistrictName string  `json:"districtName"`
	ProvinceId   *uint32 `json:"provinceId"`
}
