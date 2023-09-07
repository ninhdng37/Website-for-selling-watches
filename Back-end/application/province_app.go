package application

import (
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type provinceApp struct {
	p repository.ProvinceRepository
}

var _ ProvinceAppInterface = &provinceApp{}

type ProvinceAppInterface interface {
	GetAllProvince() ([]entity.Province, error)
	CreateProvince(provinceName string) error
	UpdateProvince(province *entity.Province) error
	DeleteProvince(provinceID *uint32) error
}

func (p *provinceApp) GetAllProvince() ([]entity.Province, error) {
	return p.p.GetAllProvince()
}

func (p *provinceApp) CreateProvince(provinceName string) error {
	return p.p.CreateProvince(provinceName)
}

func (p *provinceApp) UpdateProvince(province *entity.Province) error {
	return p.p.UpdateProvince(province)
}

func (p *provinceApp) DeleteProvince(provinceID *uint32) error {
	return p.p.DeleteProvince(provinceID)
}
