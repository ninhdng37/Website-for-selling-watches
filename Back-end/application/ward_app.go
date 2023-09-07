package application

import (
	"watchWebsite/domain/entity"
	"watchWebsite/domain/repository"
)

type wardApp struct {
	w repository.WardRepository
}

var _ WardAppInterface = &wardApp{}

type WardAppInterface interface {
	GetAllWard() ([]entity.Ward, error)
}

func (w *wardApp) GetAllWard() ([]entity.Ward, error) {
	return w.w.GetAllWard()
}
