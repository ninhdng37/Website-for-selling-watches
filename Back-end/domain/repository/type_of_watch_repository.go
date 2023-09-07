package repository

import "watchWebsite/domain/entity"

type TypeOfWatchRepository interface {
	GetAllTypesOfWatch() ([]*entity.TypeOfWatch, error)
	GetTypeOfWatchByName(typeOfWatchName string) (*entity.TypeOfWatch, error)
	CreateTypeOfWatch(typeOfWatchName string) error
	UpdateTypeOfWatch(typeOfWatch *entity.TypeOfWatch) error
	DeleteTypeOfWatch(TypeOfWatchID *uint32) error
}
