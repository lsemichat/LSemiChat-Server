package service

import (
	"app/api/domain/entity"
	"app/api/domain/repository"
	"time"

	"github.com/pkg/errors"
)

type MessageService interface {
	New(message string, grade int, author *entity.User, thread *entity.Thread) (*entity.Message, error)
	GetByID(id string) (*entity.Message, error)
	GetByThreadID(threadID string) ([]*entity.Message, error)
	AddFavorite(messageID, userUUID string) error
}

type messageService struct {
	messageRepository repository.MessageRepository
}

func NewMessageService(mr repository.MessageRepository) MessageService {
	return &messageService{
		messageRepository: mr,
	}
}

func (ms *messageService) New(message string, grade int, author *entity.User, thread *entity.Thread) (*entity.Message, error) {
	id, err := GenerateUUID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate id")
	}
	now := time.Now()
	msg := &entity.Message{
		ID:        id,
		Message:   message,
		Grade:     grade,
		CreatedAt: &now,
		Author:    author,
		Thread:    thread,
	}
	if err = ms.messageRepository.Create(msg); err != nil {
		return nil, errors.Wrap(err, "failed to create message")
	}
	return msg, nil
}

func (ms *messageService) GetByID(id string) (*entity.Message, error) {
	message, err := ms.messageRepository.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get message")
	}
	return message, nil
}

func (ms *messageService) GetByThreadID(threadID string) ([]*entity.Message, error) {
	messages, err := ms.messageRepository.GetByThreadID(threadID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get messages")
	}
	return messages, nil
}

func (ms *messageService) AddFavorite(messageID, userUUID string) error {
	id, err := GenerateUUID()
	if err != nil {
		return errors.Wrap(err, "failed to generate id")
	}
	if err = ms.messageRepository.AddFavorite(id, messageID, userUUID); err != nil {
		return errors.Wrap(err, "failed to add relation")
	}
	return nil
}
