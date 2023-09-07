package persistence

import (
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type DistrictRepo struct {
	db *Database
}

func NewDistrictRepository(db *Database) *DistrictRepo {
	return &DistrictRepo{db}
}

var _ repository.DistrictRepository = &DistrictRepo{}

func (d *DistrictRepo) GetAllDistrict() ([]entity.District, error) {
	rows, err := d.db.db.Query(`	SELECT *
							FROM districts`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var districts []entity.District

	for rows.Next() {
		district := entity.District{}
		if err := rows.Scan(
			&district.DistrictId,
			&district.DistrictName,
			&district.ProvinceId); err != nil {
			return nil, err
		}
		districts = append(districts, district)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return districts, nil
}
