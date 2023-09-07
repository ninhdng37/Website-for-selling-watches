package application

import (
	"database/sql"
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type watchApp struct {
	wr repository.WatchRepository
}

var _ WatchAppInterface = &watchApp{}

type WatchAppInterface interface {
	GetAllWatch() ([]entity.Watch, error)
	GetWatchByName(name string) (*entity.Watch, error)
	GetWatchByNameRelative(name string) ([]*entity.Watch, error)
	UpdateWatch(tx *sql.Tx, watch *entity.Watch) error
	UpdateWatchCus(*entity.Watch) error
	CreateWatch(*entity.Watch) (*uint32, error)
	DeleteWatch(watchID *uint32) error
}

func (w *watchApp) GetAllWatch() ([]entity.Watch, error) {
	return w.wr.GetAllWatch()
}

func (w *watchApp) GetWatchByName(name string) (*entity.Watch, error) {
	return w.wr.GetWatchByName(name)
}

func (w *watchApp) GetWatchByNameRelative(name string) ([]*entity.Watch, error) {
	return w.wr.GetWatchByNameRelative(name)
}

func (w *watchApp) UpdateWatch(tx *sql.Tx, watch *entity.Watch) error {
	return w.wr.UpdateWatch(tx, watch)
}

func (w *watchApp) UpdateWatchCus(watch *entity.Watch) error {
	return w.wr.UpdateWatchCus(watch)
}

func (w *watchApp) CreateWatch(watch *entity.Watch) (*uint32, error) {
	return w.wr.CreateWatch(watch)
}

func (w *watchApp) DeleteWatch(watchID *uint32) error {
	return w.wr.DeleteWatch(watchID)
}
