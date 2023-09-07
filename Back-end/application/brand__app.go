package application

import (
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type brandApp struct {
	b repository.BrandRepository
}

var _ BrandAppInterface = &brandApp{}

type BrandAppInterface interface {
	GetAllBrands() ([]*entity.Brand, error)
	GetBrandByID(brandID *uint32) (*entity.Brand, error)
	CreateBrand(brandName string) error
	UpdateBrand(brand *entity.Brand) error
	DeleteBrand(brandID *uint32) error
}

func (b *brandApp) GetAllBrands() ([]*entity.Brand, error) {
	return b.b.GetAllBrands()
}

func (b *brandApp) GetBrandByID(brandID *uint32) (*entity.Brand, error) {
	return b.b.GetBrandByID(brandID)
}
func (b *brandApp) CreateBrand(brandName string) error {
	return b.b.CreateBrand(brandName)
}

func (b *brandApp) UpdateBrand(brand *entity.Brand) error {
	return b.b.UpdateBrand(brand)
}

func (b *brandApp) DeleteBrand(brandID *uint32) error {
	return b.b.DeleteBrand(brandID)
}
