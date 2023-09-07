package application

import (
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type districtApp struct {
	d repository.DistrictRepository
}

var _ DistrictAppInterface = &districtApp{}

type DistrictAppInterface interface {
	GetAllDistrict() ([]entity.District, error)
}

func (d *districtApp) GetAllDistrict() ([]entity.District, error) {
	return d.d.GetAllDistrict()
}
