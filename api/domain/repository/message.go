package repository

import "app/api/domain/entity"

type MessageRepository interface {
	Create(message *entity.Message) error
	GetByThreadID(threadID string) ([]*entity.Message, error)
	GetByID(id string) (*entity.Message, error)
	AddFavorite(id, messageID, userUUID string) error
}
