package service

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	PasswordEncrypt(password string) (string, error)
	VerifyPassword(hash, password string) error
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (ah *authService) PasswordEncrypt(password string) (string, error) {
	if len(password) > 70 {
		// NOTE: 使っているパッケージの性質上、72文字以上のパスワードだと認証漏れするため
		return "", errors.New("password is less 70 chatacters")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (ah *authService) VerifyPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
