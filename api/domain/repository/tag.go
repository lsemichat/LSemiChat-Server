package repository

import "app/api/domain/entity"

type TagRepository interface {
	Create(tag *entity.Tag) error
	FindAll() ([]*entity.Tag, error)
	FindByID(id string) (*entity.Tag, error)
	FindByCategoryID(id string) ([]*entity.Tag, error)
	FindByUserUUID(id string) ([]*entity.Tag, error)
	FindByThreadID(id string) ([]*entity.Tag, error)
	FindByTagAndCategoryID(tagValue, categoryID string) (*entity.Tag, error)
	FindByTagNames(tagNames []string) ([]*entity.Tag, error)
	AddToUser(id, tagID, userUUID string) error
	RemoveFromUser(tagID, userUUID string) error
	AddToThread(id, tagID, threadID string) error
	RemoveFromThread(tagID, threadID string) error
}
