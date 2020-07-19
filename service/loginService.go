package service

import (
	"github.com/seinyan/go-rest-api/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	GeneratePasswordHash(password string) (string, error)
	compareHashAndPassword(password string, passwordHash string) error
	IsAuthenticated(user models.User) bool
}

type loginService struct {}

func NewLoginService() LoginService {
	return &loginService{}
}

func (service loginService) GeneratePasswordHash(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}

func (service loginService) IsAuthenticated(user models.User) bool {
	if err := service.compareHashAndPassword(user.Password, user.PasswordHash); err != nil {
		return false
	}
	return true
}

func (service loginService) compareHashAndPassword(password string, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}