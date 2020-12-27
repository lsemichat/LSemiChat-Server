package service

import (
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
	// TODO: 文字数のバリデーション
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (ah *authService) VerifyPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
