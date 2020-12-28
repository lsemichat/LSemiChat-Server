package interactor

import (
	"app/api/domain/entity"
	"app/api/domain/service"

	"github.com/pkg/errors"
)

type MessageInteractor interface {
	Create(message string, grade int, authorID string, threadID string) (*entity.Message, error)
	GetByID(id string) (*entity.Message, error)
	GetByThreadID(threadID string) ([]*entity.Message, error)
	AddFavorite(messageID, userID string) error
}

type messageInteractor struct {
	messageService service.MessageService
	threadService  service.ThreadService
	userService    service.UserService
}

func NewMessageInteractor(ms service.MessageService, ts service.ThreadService, us service.UserService) MessageInteractor {
	return &messageInteractor{
		messageService: ms,
		threadService:  ts,
		userService:    us,
	}
}

func (mi *messageInteractor) Create(message string, grade int, authorID string, threadID string) (*entity.Message, error) {
	thread, err := mi.threadService.GetByID(threadID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find thread")
	}
	author, err := mi.userService.GetByUserID(authorID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find user")
	}
	msg, err := mi.messageService.New(message, grade, author, thread)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create message")
	}
	msg.Author = author
	msg.Thread = thread
	return msg, nil
}

func (mi *messageInteractor) GetByID(id string) (*entity.Message, error) {
	message, err := mi.messageService.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get message")
	}

	author, err := mi.userService.GetByID(message.Author.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get author")
	}

	thread, err := mi.threadService.GetByID(message.Thread.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get thread")
	}
	message.Author = author
	message.Thread = thread
	return message, nil
}

func (mi *messageInteractor) GetByThreadID(threadID string) ([]*entity.Message, error) {
	messages, err := mi.messageService.GetByThreadID(threadID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get messages")
	}
	return messages, nil
}

func (mi *messageInteractor) AddFavorite(messageID, userID string) error {
	user, err := mi.userService.GetByUserID(userID)
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}
	err = mi.messageService.AddFavorite(messageID, user.ID)
	if err != nil {
		return errors.Wrap(err, "failed to add favorite")
	}
	return nil
}
