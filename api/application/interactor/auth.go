package interactor

import (
	"app/api/domain/service"

	"github.com/pkg/errors"
)

type AuthInteractor interface {
	Login(userID, password string) error
}

type authInteractor struct {
	AuthService service.AuthService
	UserService service.UserService
}

func NewAuthInteractor(as service.AuthService, us service.UserService) AuthInteractor {
	return &authInteractor{
		AuthService: as,
		UserService: us,
	}
}

func (ai *authInteractor) Login(userID, password string) error {
	user, err := ai.UserService.GetByUserID(userID)
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}

	err = ai.AuthService.VerifyPassword(user.Password, password)
	if err != nil {
		return errors.Wrap(err, "failed to verify password")
	}

	return nil
}
