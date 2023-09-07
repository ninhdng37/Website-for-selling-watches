package persistence

import (
	"database/sql"
	"errors"
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type WatchRepo struct {
	db *Database
}

func NewWatchRepository(db *Database) *WatchRepo {
	return &WatchRepo{db}
}

var _ repository.WatchRepository = &WatchRepo{}

func (w *WatchRepo) GetAllWatch() ([]entity.Watch, error) {
	rows, err := w.db.db.Query(`	SELECT *
							FROM watches
							WHERE status != 2`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var watches []entity.Watch

	for rows.Next() {
		watch := entity.Watch{}
		if err := rows.Scan(
			&watch.WatchId,
			&watch.WatchName,
			&watch.Price,
			&watch.Image,
			&watch.Status,
			&watch.Quantity,
			&watch.BrandId,
			&watch.TypeOfWatchId); err != nil {
			return nil, err
		}
		watches = append(watches, watch)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return watches, nil
}

func (w *WatchRepo) GetWatchByName(name string) (*entity.Watch, error) {
	rows := w.db.db.QueryRow(`SELECT *
	FROM watches
	WHERE watch_name = $1`, name)

	watch := &entity.Watch{}
	err := rows.Scan(
		&watch.WatchId,
		&watch.WatchName,
		&watch.Price,
		&watch.Image,
		&watch.Status,
		&watch.Quantity,
		&watch.BrandId,
		&watch.TypeOfWatchId,
	)

	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return watch, nil
}

func (w *WatchRepo) GetWatchByNameRelative(name string) ([]*entity.Watch, error) {
	name = "%" + name + "%"
	// fmt.Println(name)
	rows, err := w.db.db.Query(`	SELECT *
							FROM watches
							WHERE watch_name LIKE $1`, name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var watches []*entity.Watch

	for rows.Next() {
		watch := &entity.Watch{}
		if err := rows.Scan(
			&watch.WatchId,
			&watch.WatchName,
			&watch.Price,
			&watch.Image,
			&watch.Status,
			&watch.Quantity,
			&watch.BrandId,
			&watch.TypeOfWatchId); err != nil {
			return nil, err
		}
		watches = append(watches, watch)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return watches, nil
}

func (w *WatchRepo) UpdateWatch(tx *sql.Tx, watch *entity.Watch) error {
	// tx, err := w.db.Begin()
	// if err != nil {
	// 	return err
	// }
	// defer tx.Rollback()
	query := `	
	UPDATE watches
	SET 
	watch_name = CASE WHEN $1 = '-1' THEN watch_name ELSE $1 END,
	price = CASE WHEN $2 = -1 THEN price ELSE $2 END,
	image = CASE WHEN $3 = '-1' THEN image ELSE $3 END,
	status = CASE WHEN $4 = -1 THEN status ELSE $4 END,
	quantity = CASE WHEN $5 = -1 THEN quantity ELSE $5 END,
	brand_id = CASE WHEN $6 = 0 THEN brand_id ELSE $6 END,
	type_of_watch_id = CASE WHEN $7 = 0 THEN type_of_watch_id ELSE $7 END
	WHERE watch_id = $8`
	_, err := tx.Exec(
		query,
		watch.WatchName,
		watch.Price,
		watch.Image,
		watch.Status,
		watch.Quantity,
		watch.BrandId,
		watch.TypeOfWatchId,
		watch.WatchId)

	if err != nil {
		return err
	}

	// tx.Commit()
	return nil
}

func (w *WatchRepo) CreateWatch(watch *entity.Watch) (*uint32, error) {

	tx, err := w.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()
	query := `	
	INSERT INTO watches (
	watch_name,
	price,
	quantity,
	brand_id,
	type_of_watch_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING watch_id;`

	var watchID uint32

	err = tx.QueryRow(
		query,
		watch.WatchName,
		watch.Price,
		watch.Quantity,
		watch.BrandId,
		watch.TypeOfWatchId).Scan(&watchID)

	if err != nil {
		return nil, err
	}
	tx.Commit()
	return &watchID, nil
}

func (w *WatchRepo) UpdateWatchCus(watch *entity.Watch) error {
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `	
	UPDATE watches
	SET 
	watch_name = CASE WHEN $1 = '-1' THEN watch_name ELSE $1 END,
	price = CASE WHEN $2 = 0 THEN price ELSE $2 END,
	image = CASE WHEN $3 = '-1' THEN image ELSE $3 END,
	status = CASE WHEN $4 = -1 THEN status ELSE $4 END,
	quantity = CASE WHEN $5 = -1 THEN quantity ELSE $5 END,
	brand_id = CASE WHEN $6 = 0 THEN brand_id ELSE $6 END,
	type_of_watch_id = CASE WHEN $7 = 0 THEN type_of_watch_id ELSE $7 END
	WHERE watch_id = $8`
	_, err = tx.Exec(
		query,
		watch.WatchName,
		watch.Price,
		watch.Image,
		watch.Status,
		watch.Quantity,
		watch.BrandId,
		watch.TypeOfWatchId,
		watch.WatchId)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (w *WatchRepo) DeleteWatch(watchID *uint32) error {
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var count int
	query := "SELECT COUNT(*) FROM order_details WHERE watch_id = $1"
	err = tx.QueryRow(query, watchID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("can not delete this watch because this watch was bought")
	}
	_, err = w.db.db.Exec("DELETE FROM watches WHERE watch_id = $1", watchID)

	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
