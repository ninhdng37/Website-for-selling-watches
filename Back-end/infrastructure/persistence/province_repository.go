package persistence

import (
	"database/sql"
	"errors"
	"fmt"
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type ProvinceRepo struct {
	db *Database
}

func NewProvinceRepository(db *Database) *ProvinceRepo {
	return &ProvinceRepo{db}
}

var _ repository.ProvinceRepository = &ProvinceRepo{}

func (p *ProvinceRepo) GetAllProvince() ([]entity.Province, error) {
	rows, err := p.db.db.Query(`	SELECT *
							FROM provinces`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var provinces []entity.Province

	for rows.Next() {
		province := entity.Province{}
		if err := rows.Scan(
			&province.ProvinceId,
			&province.ProvinceName); err != nil {
			return nil, err
		}
		provinces = append(provinces, province)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return provinces, nil
}

func (p *ProvinceRepo) GetProvinceByName(provinceName string) (*entity.Province, error) {

	row := p.db.db.QueryRow(`SELECT *
	FROM provinces
	WHERE province_name = $1`, provinceName)

	province := &entity.Province{}
	err := row.Scan(
		&province.ProvinceId,
		&province.ProvinceName)
	if err != nil {
		return nil, err
	}

	if err = row.Err(); err != nil {
		return nil, err
	}
	return province, nil
}

func (p *ProvinceRepo) CreateProvince(provinceName string) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	province, err := p.GetProvinceByName(provinceName)
	if province != nil {
		return errors.New("this province name existed")
	}

	if err != nil && err != sql.ErrNoRows {
		fmt.Println(10)
		fmt.Println(err)
		return err
	}

	defer tx.Rollback()
	query := `	
	INSERT INTO provinces (
				province_name)
	VALUES ($1)`

	_, err = tx.Exec(
		query,
		provinceName)

	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (p *ProvinceRepo) UpdateProvince(province *entity.Province) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `	
	UPDATE provinces
	SET 
	province_name = CASE WHEN $1 = '-1' THEN province_name ELSE $1 END
	WHERE province_id = $2`
	_, err = tx.Exec(
		query,
		province.ProvinceName,
		province.ProvinceId)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (p *ProvinceRepo) DeleteProvince(provinceID *uint32) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var count int
	query := "SELECT COUNT(*) FROM districts WHERE province_id = $1"
	err = tx.QueryRow(query, provinceID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("can not delete province because this province was assigned for another district")
	}
	_, err = p.db.db.Exec("DELETE FROM provinces WHERE province_id = $1", provinceID)

	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
