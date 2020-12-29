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
	AddFollow(userID, followedUserID string) error
	DeleteFollow(userID, followedUserID string) error
	GetFollowers(id string) ([]*entity.User, error)
}

type userInteractor struct {
	userService     service.UserService
	authService     service.AuthService
	tagService      service.TagService
	categoryService service.CategoryService
}

func NewUserInteractor(us service.UserService, as service.AuthService, ts service.TagService, cs service.CategoryService) UserInteractor {
	return &userInteractor{
		userService:     us,
		authService:     as,
		tagService:      ts,
		categoryService: cs,
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
	tags, err := ui.tagService.GetByUserUUID(newUser.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tags")
	}
	newUser.Tags = AddCategoryToTag(tags, ui.categoryService)
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
	tags, err := ui.tagService.GetByUserUUID(newUser.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tags")
	}
	newUser.Tags = AddCategoryToTag(tags, ui.categoryService)
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
	tags, err := ui.tagService.GetByUserUUID(newUser.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tags")
	}
	newUser.Tags = AddCategoryToTag(tags, ui.categoryService)
	return newUser, nil
}

func (ui *userInteractor) GetByID(id string) (*entity.User, error) {
	user, err := ui.userService.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}
	tags, err := ui.tagService.GetByUserUUID(user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tags")
	}
	user.Tags = AddCategoryToTag(tags, ui.categoryService)
	return user, nil
}

func (ui *userInteractor) GetByUserID(userID string) (*entity.User, error) {
	user, err := ui.userService.GetByUserID(userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}
	tags, err := ui.tagService.GetByUserUUID(user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tags")
	}
	user.Tags = AddCategoryToTag(tags, ui.categoryService)
	return user, nil
}

func (ui *userInteractor) GetByMail(mail string) (*entity.User, error) {
	user, err := ui.userService.GetByMail(mail)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}
	tags, err := ui.tagService.GetByUserUUID(user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tags")
	}
	user.Tags = AddCategoryToTag(tags, ui.categoryService)
	return user, nil
}

func (ui *userInteractor) GetAll() ([]*entity.User, error) {
	users, err := ui.userService.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get users")
	}
	result := make([]*entity.User, 0, len(users))
	for _, user := range users {
		tags, _ := ui.tagService.GetByUserUUID(user.ID)
		user.Tags = AddCategoryToTag(tags, ui.categoryService)
		result = append(result, user)
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
	result := make([]*entity.User, 0, len(users))
	for _, user := range users {
		tags, _ := ui.tagService.GetByUserUUID(user.ID)
		user.Tags = AddCategoryToTag(tags, ui.categoryService)
		result = append(result, user)
	}
	return users, nil
}

func (ui *userInteractor) AddFollow(userID, followedUserID string) error {
	user, err := ui.userService.GetByUserID(userID)
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}
	err = ui.userService.AddFollow(user.ID, followedUserID)
	if err != nil {
		return errors.Wrap(err, "failed to add follow")
	}
	return nil
}

func (ui *userInteractor) DeleteFollow(userID, followedUserID string) error {
	user, err := ui.userService.GetByUserID(userID)
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}
	err = ui.userService.DeleteFollow(user.ID, followedUserID)
	if err != nil {
		return errors.Wrap(err, "failed to delete follow")
	}
	return nil
}

func (ui *userInteractor) GetFollowers(id string) ([]*entity.User, error) {
	users, err := ui.userService.GetFollowers(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get followers")
	}
	result := make([]*entity.User, 0, len(users))
	for _, user := range users {
		tags, _ := ui.tagService.GetByUserUUID(user.ID)
		user.Tags = AddCategoryToTag(tags, ui.categoryService)
		result = append(result, user)
	}
	return users, nil
}
