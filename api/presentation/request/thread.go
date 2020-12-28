package request

import (
	"app/api/application/interactor"

	"github.com/pkg/errors"
)

type CreateThreadRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LimitUsers  int    `json:"limit_users"`
	IsPublic    int    `json:"is_public"`
}

func (r *CreateThreadRequest) Validation() error {
	if r.Name == "" {
		return errors.New("required filed is empty")
	}
	if r.IsPublic != 0 && r.IsPublic != 1 {
		return errors.New("is_public allow 0 or 1")
	}
	if r.LimitUsers < 0 {
		return errors.New("limit_users allow cant minas")
	}
	return nil
}

type UpdateThreadRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LimitUsers  int    `json:"limit_users"`
	IsPublic    int    `json:"is_public"`
}

func (r *UpdateThreadRequest) Validation(ti interactor.ThreadInteractor, threadID, requestUserID string) error {
	if r.Name == "" {
		return errors.New("required filed is empty")
	}
	if r.IsPublic != 0 && r.IsPublic != 1 {
		return errors.New("is_public allow 0 or 1")
	}

	if r.LimitUsers < 1 {
		return errors.New("limit_users allow more 1")
	}

	thread, err := ti.GetByID(threadID)
	if err != nil {
		return errors.New("thread is not found")
	}
	if thread.Author.UserID != requestUserID {
		return errors.New("no authorized")
	}

	return nil
}
