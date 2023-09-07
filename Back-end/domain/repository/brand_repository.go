package repository

import "watchWebsite/domain/entity"

type BrandRepository interface {
	GetAllBrands() ([]*entity.Brand, error)
	GetBrandByID(brandID *uint32) (*entity.Brand, error)
	GetBrandByName(brandName string) (*entity.Brand, error)
	CreateBrand(brandName string) error
	UpdateBrand(brand *entity.Brand) error
	DeleteBrand(brandID *uint32) error
}
