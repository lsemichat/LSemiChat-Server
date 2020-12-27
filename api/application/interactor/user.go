package interactor

import (
	"app/api/domain/entity"
	"app/api/domain/service"

	"github.com/pkg/errors"
)

type UserInteractor interface {
	Create(userID, name, mail, image, profile, password string) (*entity.User, error)
	UpdateProfile(userID, name, mail, image, profile string) (*entity.User, error)
	UpdateUserID(userID, newUserID string) (*entity.User, error)
	UpdatePassword(userID, password string) (*entity.User, error)
	GetByID(id string) (*entity.User, error)
	GetByUserID(userID string) (*entity.User, error)
	GetByMail(mail string) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Delete(userID string) error
	GetFollows(id string) ([]*entity.User, error)
	GetFollowers(id string) ([]*entity.User, error)
}

type userInteractor struct {
	userService service.UserService
	authService service.AuthService
}

func NewUserInteractor(us service.UserService, as service.AuthService) UserInteractor {
	return &userInteractor{
		userService: us,
		authService: as,
	}
}

func (ui *userInteractor) Create(userID, name, mail, image, profile, password string) (*entity.User, error) {
	hash, err := ui.authService.PasswordEncrypt(password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate password")
	}
	user, err := ui.userService.New(userID, name, mail, image, profile, hash, 0)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new entity")
	}
	return user, nil
}

func (ui *userInteractor) UpdateProfile(userID, name, mail, image, profile string) (*entity.User, error) {
	user, err := ui.userService.GetByUserID(userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}
	newUser, err := ui.userService.UpdateProfile(user, name, mail, image, profile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update profile")
	}
	return newUser, nil
}

func (ui *userInteractor) UpdateUserID(userID, newUserID string) (*entity.User, error) {
	user, err := ui.userService.GetByUserID(userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}
	newUser, err := ui.userService.UpdateUserID(user, newUserID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update userID")
	}
	return newUser, nil
}

func (ui *userInteractor) UpdatePassword(userID, password string) (*entity.User, error) {
	hash, err := ui.authService.PasswordEncrypt(password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate password")
	}
	user, err := ui.userService.GetByUserID(userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}
	newUser, err := ui.userService.UpdatePassword(user, hash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update password")
	}
	return newUser, nil
}

func (ui *userInteractor) GetByID(id string) (*entity.User, error) {
	user, err := ui.userService.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}
	return user, nil
}

func (ui *userInteractor) GetByUserID(userID string) (*entity.User, error) {
	user, err := ui.userService.GetByUserID(userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}
	return user, nil
}

func (ui *userInteractor) GetByMail(mail string) (*entity.User, error) {
	user, err := ui.userService.GetByMail(mail)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}
	return user, nil
}

func (ui *userInteractor) GetAll() ([]*entity.User, error) {
	users, err := ui.userService.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get users")
	}
	return users, nil
}

func (ui *userInteractor) Delete(userID string) error {
	user, err := ui.userService.GetByUserID(userID)
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}
	err = ui.userService.Delete(user.ID)
	if err != nil {
		return errors.Wrap(err, "failed to delete")
	}
	return nil
}

func (ui *userInteractor) GetFollows(id string) ([]*entity.User, error) {
	users, err := ui.userService.GetFollows(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get follows")
	}
	return users, nil
}

func (ui *userInteractor) GetFollowers(id string) ([]*entity.User, error) {
	users, err := ui.userService.GetFollowers(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get followers")
	}
	return users, nil
}
