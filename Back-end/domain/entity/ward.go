package entity

type Ward struct {
	WardId     *uint32 `json:"wardId"`
	WardName   string  `json:"wardName"`
	DistrictId *uint32 `json:"districtId"`
}
