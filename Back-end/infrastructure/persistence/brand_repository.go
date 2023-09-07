package persistence

import (
	"database/sql"
	"errors"
	"fmt"
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type BrandRepo struct {
	db *Database
}

func NewBrandRepository(db *Database) *BrandRepo {
	return &BrandRepo{db}
}

var _ repository.BrandRepository = &BrandRepo{}

func (b *BrandRepo) GetAllBrands() ([]*entity.Brand, error) {
	rows, err := b.db.db.Query(`	SELECT *
									FROM brands
									ORDER BY brand_id`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var brands []*entity.Brand

	for rows.Next() {
		brand := &entity.Brand{}
		if err := rows.Scan(
			&brand.BrandID,
			&brand.BrandName); err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return brands, nil
}

func (b *BrandRepo) GetBrandByID(brandID *uint32) (*entity.Brand, error) {
	rows := b.db.db.QueryRow(`SELECT *
	FROM brands
	WHERE brand_id = $1`, brandID)

	brand := &entity.Brand{}
	err := rows.Scan(
		&brand.BrandID,
		&brand.BrandName)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return brand, nil
}

func (b *BrandRepo) GetBrandByName(brandName string) (*entity.Brand, error) {
	rows := b.db.db.QueryRow(`SELECT *
	FROM brands
	WHERE brand_name = $1`, brandName)

	brand := &entity.Brand{}
	err := rows.Scan(
		&brand.BrandID,
		&brand.BrandName)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return brand, nil
}

func (b *BrandRepo) CreateBrand(brandName string) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}

	brand, err := b.GetBrandByName(brandName)
	if brand != nil {
		return errors.New("this brand name existed")
	}

	if err != nil && err != sql.ErrNoRows {
		fmt.Println(10)
		fmt.Println(err)
		return err
	}

	defer tx.Rollback()
	query := `	
	INSERT INTO brands (
	brand_name)
	VALUES ($1)`

	_, err = tx.Exec(
		query,
		brandName)

	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (b *BrandRepo) UpdateBrand(brand *entity.Brand) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `	
	UPDATE brands
	SET 
	brand_name = CASE WHEN $1 = '-1' THEN brand_name ELSE $1 END
	WHERE brand_id = $2`
	_, err = tx.Exec(
		query,
		brand.BrandName,
		brand.BrandID)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (b *BrandRepo) DeleteBrand(brandID *uint32) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var count int
	query := "SELECT COUNT(*) FROM watches WHERE brand_id = $1"
	err = tx.QueryRow(query, brandID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("can not delete brand because this brand was assigned for watches")
	}
	_, err = b.db.db.Exec("DELETE FROM brands WHERE brand_id = $1", brandID)

	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
