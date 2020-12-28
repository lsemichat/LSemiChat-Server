package request

import (
	"app/api/application/interactor"

	"github.com/pkg/errors"
)

type CreateTagRequest struct {
	Tag        string `json:"tag"`
	CategoryID string `json:"category_id"`
}

func (r *CreateTagRequest) Validation(ci interactor.CategoryInteractor) error {
	if r.Tag == "" || r.CategoryID == "" {
		return errors.New("required field is empty")
	}

	// TODO: GetByCategoryID
	_, err := ci.GetByID(r.CategoryID)
	if err != nil {
		return errors.New("can't find category")
	}
	return nil
}
