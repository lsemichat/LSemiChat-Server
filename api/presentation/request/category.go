package request

import "github.com/pkg/errors"

type CreateCategoryRequest struct {
	Category string `json:"category"`
}

func (r *CreateCategoryRequest) Validate() error {
	if r.Category == "" {
		return errors.New("required field is empty")
	}
	return nil
}

type UpdateCategoryRequest struct {
	Category string `json:"category"`
}

func (r *UpdateCategoryRequest) Validate() error {
	if r.Category == "" {
		return errors.New("required field is empty")
	}
	return nil
}
