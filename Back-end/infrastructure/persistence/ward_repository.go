package persistence

import (
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type WardRepo struct {
	db *Database
}

func NewWardRepository(db *Database) *WardRepo {
	return &WardRepo{db}
}

var _ repository.WardRepository = &WardRepo{}

func (w *WardRepo) GetAllWard() ([]entity.Ward, error) {
	rows, err := w.db.db.Query(`	SELECT *
							FROM wards`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var wards []entity.Ward

	for rows.Next() {
		ward := entity.Ward{}
		if err := rows.Scan(
			&ward.WardId,
			&ward.WardName,
			&ward.DistrictId); err != nil {
			return nil, err
		}
		wards = append(wards, ward)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return wards, nil
}
