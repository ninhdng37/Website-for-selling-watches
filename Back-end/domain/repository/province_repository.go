package repository

import "watchWebsite/domain/entity"

type ProvinceRepository interface {
	GetAllProvince() ([]entity.Province, error)
	GetProvinceByName(provinceName string) (*entity.Province, error)
	CreateProvince(provinceName string) error
	UpdateProvince(province *entity.Province) error
	DeleteProvince(provinceID *uint32) error
}
