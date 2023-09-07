package repository

import "watchWebsite/domain/entity"

type DistrictRepository interface {
	GetAllDistrict() ([]entity.District, error)
}
