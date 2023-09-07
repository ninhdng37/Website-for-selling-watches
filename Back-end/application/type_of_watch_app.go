package application

import (
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type typeOfWatchApp struct {
	b repository.TypeOfWatchRepository
}

var _ TypeOfWatchAppInterface = &typeOfWatchApp{}

type TypeOfWatchAppInterface interface {
	GetAllTypesOfWatch() ([]*entity.TypeOfWatch, error)
	GetTypeOfWatchByName(typeOfWatchName string) (*entity.TypeOfWatch, error)
	CreateTypeOfWatch(typeOfWatchName string) error
	UpdateTypeOfWatch(typeOfWatch *entity.TypeOfWatch) error
	DeleteTypeOfWatch(typeOfWatchID *uint32) error
}

func (b *typeOfWatchApp) GetAllTypesOfWatch() ([]*entity.TypeOfWatch, error) {
	return b.b.GetAllTypesOfWatch()
}

func (b *typeOfWatchApp) GetTypeOfWatchByName(typeOfWatchName string) (*entity.TypeOfWatch, error) {
	return b.b.GetTypeOfWatchByName(typeOfWatchName)
}

func (b *typeOfWatchApp) CreateTypeOfWatch(typeOfWatchName string) error {
	return b.b.CreateTypeOfWatch(typeOfWatchName)
}

func (b *typeOfWatchApp) UpdateTypeOfWatch(typeOfWatch *entity.TypeOfWatch) error {
	return b.b.UpdateTypeOfWatch(typeOfWatch)
}

func (b *typeOfWatchApp) DeleteTypeOfWatch(typeOfWatchID *uint32) error {
	return b.b.DeleteTypeOfWatch(typeOfWatchID)
}
