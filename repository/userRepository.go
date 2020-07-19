package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/seinyan/go-rest-api/models"
)

type UserRepository interface {
	Get(id uint64) (models.User, error)
	Create(item models.User) (models.User, error)
	Update(item models.User) error
	Delete(item models.User) error
	Register(item models.User) (models.User, error)
	GetByUsername(username string) (models.User, error)
}

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (repo userRepository) Get(id uint64) (models.User, error) {
	var item models.User
	err := repo.conn.First(&item, id).Error
	return item, err
}


func (repo userRepository) Create(item models.User) (models.User, error) {
	repo.conn.NewRecord(item)
	err := repo.conn.Create(&item).Error
	return item, err
}

func (repo userRepository) Update(item models.User) error {
	err := repo.conn.Save(&item).Error
	return err
}

func (repo userRepository) Delete(item models.User) error {
	err := repo.conn.Delete(&item).Error
	return err
}

func (repo userRepository) GetByUsername(username string) (models.User, error) {
	var item models.User
	err := repo.conn.Where("email = ?", username).First(&item).Error
	return item, err
}



func (repo userRepository) Register(item models.User) (models.User, error) {
	repo.conn.NewRecord(item)
	err := repo.conn.Create(&item).Error
	return item, err
}