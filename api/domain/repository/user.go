package repository

import (
	"app/api/domain/entity"
	"time"
)

type UserRepository interface {
	Create(user *entity.User) error
	UpdateProfile(user *entity.User) error
	UpdateUserID(id, userID string, updatedAt *time.Time) error
	UpdatePassword(id, password string, updatedAt *time.Time) error
	FindAll() ([]*entity.User, error)
	FindByID(id string) (*entity.User, error)
	FindByUserID(userID string) (*entity.User, error)
	FindByMail(mail string) (*entity.User, error)
	FindByUserIDs(userIDs []string) ([]*entity.User, error)
	DeleteByID(id string) error
	FindFollows(id string) ([]*entity.User, error)
	AddFollow(id, userID, followedUserID string) error
	DeleteFollow(userID, followedUserID string) error
	FindFollowers(id string) ([]*entity.User, error)
}
