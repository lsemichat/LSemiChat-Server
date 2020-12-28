package request

import "github.com/pkg/errors"

type CreateMessageRequest struct {
	Message string `json:"message"`
	Grade   int    `json:"grade"`
}

func (r *CreateMessageRequest) Validation() error {
	if r.Message == "" {
		return errors.New("required field is empty")
	}
	if r.Grade < 1 {
		return errors.New("grande don't allow minas")
	}
	return nil
}
