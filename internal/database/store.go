package database

import (
	"github.com/jinzhu/gorm"
	"github.com/seinyan/go-rest-api/internal/repository"
)

type Store struct {
	Conn           *gorm.DB
	UserRepository repository.UserRepository
}

func NewStore(conn *gorm.DB) *Store {
	return &Store{
		Conn:           conn,
		UserRepository: repository.NewUserRepository(conn),
	}
}