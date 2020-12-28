package repository

import (
	"app/api/domain/entity"
	"app/api/domain/repository"
	"app/api/infrastructure/database"

	"github.com/pkg/errors"
)

type messageRepository struct {
	sqlHandler database.SQLHandler
}

func NewMessageRepository(sh database.SQLHandler) repository.MessageRepository {
	return &messageRepository{
		sqlHandler: sh,
	}
}

func (mr *messageRepository) Create(message *entity.Message) error {
	_, err := mr.sqlHandler.Exec(`
		INSERT INTO messages(id, message, grade, created_at, thread_id, user_id)
		VALUES (?, ?, ?, ?, ?, ?)
	`,
		message.ID,
		message.Message,
		message.Grade,
		message.CreatedAt,
		message.Thread.ID,
		message.Author.ID,
	)
	if err != nil {
		return errors.Wrap(err, "failed to insert db")
	}
	return nil
}

func (mr *messageRepository) GetByID(id string) (*entity.Message, error) {
	row := mr.sqlHandler.QueryRow(`
		SELECT id, message, grade, created_at, thread_id, user_id
		FROM messages
		WHERE id=?
	`, id)
	var message entity.Message
	var user entity.User
	var thread entity.Thread
	if err := row.Scan(&message.ID, &message.Message, &message.Grade, &message.CreatedAt, &thread.ID, &user.ID); err != nil {
		return nil, errors.Wrap(err, "failed to scan")
	}
	message.Author = &user
	message.Thread = &thread
	return &message, nil

}

func (mr *messageRepository) GetByThreadID(threadID string) ([]*entity.Message, error) {
	rows, err := mr.sqlHandler.Query(`
		SELECT id, message, grade, created_at, thread_id, user_id
		FROM messages
		WHERE thread_id=?
	`, threadID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select")
	}
	var messages []*entity.Message
	for rows.Next() {
		var message entity.Message
		var user entity.User
		var thread entity.Thread
		if err = rows.Scan(&message.ID, &message.Message, &message.Grade, &message.CreatedAt, &thread.ID, &user.ID); err != nil {
			if rows.CheckNoRows(err) {
				return nil, nil
			}
			return nil, errors.Wrap(err, "failed to scan")
		}
		message.Author = &user
		message.Thread = &thread
		messages = append(messages, &message)
	}
	return messages, nil

}

func (mr *messageRepository) AddFavorite(id, messageID, userUUID string) error {
	_, err := mr.sqlHandler.Exec(`
		INSERT INTO users_favorites(id, user_id, message_id)
		VALUES (?, ?, ?)
	`, id, userUUID, messageID)
	if err != nil {
		return errors.Wrap(err, "failed to insert db")
	}
	return nil
}
