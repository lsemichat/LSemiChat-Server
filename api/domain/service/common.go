package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func GenerateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate id")
	}
	return id.String(), nil
}
