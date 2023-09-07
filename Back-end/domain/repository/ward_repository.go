package repository

import "watchWebsite/domain/entity"

type WardRepository interface {
	GetAllWard() ([]entity.Ward, error)
}
