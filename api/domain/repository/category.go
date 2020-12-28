package repository

import "app/api/domain/entity"

type CategoryRepository interface {
	GetAll() ([]*entity.Category, error)
	GetByID(id string) (*entity.Category, error)
	Create(category *entity.Category) error
	Update(category *entity.Category) error
	Delete(id string) error
}
