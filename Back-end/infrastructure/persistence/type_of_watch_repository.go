package persistence

import (
	"database/sql"
	"errors"
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type TypeOfWatchRepo struct {
	db *Database
}

func NewTypeOfWatchRepository(db *Database) *TypeOfWatchRepo {
	return &TypeOfWatchRepo{db}
}

var _ repository.TypeOfWatchRepository = &TypeOfWatchRepo{}

func (b *TypeOfWatchRepo) GetAllTypesOfWatch() ([]*entity.TypeOfWatch, error) {
	rows, err := b.db.db.Query(`	SELECT *
									FROM type_of_watchs
									ORDER BY type_of_watch_id`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var typesOfWatch []*entity.TypeOfWatch

	for rows.Next() {
		typeOfWatch := &entity.TypeOfWatch{}
		if err := rows.Scan(
			&typeOfWatch.TypeOfWatchID,
			&typeOfWatch.TypeOfWatchName); err != nil {
			return nil, err
		}
		typesOfWatch = append(typesOfWatch, typeOfWatch)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return typesOfWatch, nil
}

func (b *TypeOfWatchRepo) GetTypeOfWatchByName(typeOfWatchName string) (*entity.TypeOfWatch, error) {
	rows := b.db.db.QueryRow(`SELECT *
	FROM type_of_watchs
	WHERE type_of_watch_name = $1`, typeOfWatchName)

	type_of_watch := &entity.TypeOfWatch{}
	err := rows.Scan(
		&type_of_watch.TypeOfWatchID,
		&type_of_watch.TypeOfWatchName)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return type_of_watch, nil
}

func (b *TypeOfWatchRepo) CreateTypeOfWatch(typeOfWatchName string) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}

	type_of_watch, err := b.GetTypeOfWatchByName(typeOfWatchName)
	if type_of_watch != nil {
		return errors.New("this type of watch name existed")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	defer tx.Rollback()
	query := `	
	INSERT INTO type_of_watchs (
		type_of_watch_name)
	VALUES ($1)`

	_, err = tx.Exec(
		query,
		typeOfWatchName)

	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (b *TypeOfWatchRepo) UpdateTypeOfWatch(typeOfWatch *entity.TypeOfWatch) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `	
	UPDATE type_of_watchs
	SET 
	type_of_watch_name = CASE WHEN $1 = '-1' THEN type_of_watch_name ELSE $1 END
	WHERE type_of_watch_id = $2`
	_, err = tx.Exec(
		query,
		typeOfWatch.TypeOfWatchName,
		typeOfWatch.TypeOfWatchID)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (b *TypeOfWatchRepo) DeleteTypeOfWatch(typeOfWatchID *uint32) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var count int
	query := "SELECT COUNT(*) FROM watches WHERE type_of_watch_id = $1"
	err = tx.QueryRow(query, typeOfWatchID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("can not delete type of watch because this type of watch was assigned for watches")
	}
	_, err = b.db.db.Exec("DELETE FROM type_of_watchs WHERE type_of_watch_id = $1", typeOfWatchID)

	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
