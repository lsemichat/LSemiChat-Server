package request

import (
	"app/api/application/interactor"

	"github.com/pkg/errors"
)

type CreateUserRequest struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Image    string `json:"image"`
	Profile  string `json:"profile"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) ValidateRequest(ui interactor.UserInteractor) error {

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

type UpdateProfileRequest struct {
	Name    string `json:"name"`
	Mail    string `json:"mail"`
	Image   string `json:"image"`
	Profile string `json:"profile"`
}

func (r *UpdateProfileRequest) Validate(ui interactor.UserInteractor, userID string) error {

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

type UpdateUserIDRequest struct {
	UserID string `json:"user_id"`
}

func (r *UpdateUserIDRequest) Validate(ui interactor.UserInteractor, userID string) error {
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

type UpdatePasswordRequest struct {
	Password string `json:"password"`
}

func (r *UpdatePasswordRequest) Validate() error {
	if r.Password == "" {
		return errors.New("required field is empty")
	}
	return nil
}
