package repository

import (
	"database/sql"
	"watchWebsite/domain/entity"
)

type WatchRepository interface {
	GetAllWatch() ([]entity.Watch, error)
	GetWatchByName(name string) (*entity.Watch, error)
	GetWatchByNameRelative(name string) ([]*entity.Watch, error)
	UpdateWatch(tx *sql.Tx, watch *entity.Watch) error
	UpdateWatchCus(*entity.Watch) error
	CreateWatch(*entity.Watch) (*uint32, error)
	DeleteWatch(watchID *uint32) error
}
