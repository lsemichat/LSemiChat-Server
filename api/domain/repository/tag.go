package repository

import "app/api/domain/entity"

type TagRepository interface {
	Create(tag *entity.Tag) error
	FindAll() ([]*entity.Tag, error)
	FindByID(id string) (*entity.Tag, error)
	FindByCategoryID(id string) ([]*entity.Tag, error)
	FindByUserUUID(id string) ([]*entity.Tag, error)
	FindByThreadID(id string) ([]*entity.Tag, error)
}
