package request

import (
	"app/api/application/interactor"

	"github.com/pkg/errors"
)

type UserCreateRequest struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Image    string `json:"image"`
	Profile  string `json:"profile"`
	Password string `json:"password"`
}

func (r *UserCreateRequest) ValidateRequest(ui interactor.UserInteractor) error {

	if r.UserID == "" || r.Name == "" || r.Mail == "" || r.Password == "" {
		return errors.New("require field is empty")
	}

	if user, _ := ui.GetByUserID(r.UserID); user != nil {
		return errors.New("userID is duplicate")
	}

	if user, _ := ui.GetByMail(r.Mail); user != nil {
		return errors.New("mail is duplicate")
	}

	return nil
}

type UserUpdateProfileRequest struct {
	Name    string `json:"name"`
	Mail    string `json:"mail"`
	Image   string `json:"image"`
	Profile string `json:"profile"`
}

func (r *UserUpdateProfileRequest) Validate(ui interactor.UserInteractor, userID string) error {

	if r.Name == "" || r.Mail == "" {
		return errors.New("require field is empty")
	}

	if user, _ := ui.GetByMail(r.Mail); user != nil {
		if user.UserID != userID {
			return errors.New("mail is duplicate")
		}
	}

	return nil
}

type UserUpdateUserIDRequest struct {
	UserID string `json:"user_id"`
}

func (r *UserUpdateUserIDRequest) Validate(ui interactor.UserInteractor, userID string) error {
	if r.UserID == "" {
		return errors.New("required field is empty")
	}
	if user, _ := ui.GetByUserID(r.UserID); user != nil {
		if user.UserID != userID {
			return errors.New("userID is duplicate")
		}
	}
	return nil
}

type UserUpdatePasswordRequest struct {
	Password string `json:"password"`
}

func (r *UserUpdatePasswordRequest) Validate() error {
	if r.Password == "" {
		return errors.New("required field is empty")
	}
	return nil
}
