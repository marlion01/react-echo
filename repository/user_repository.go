package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
type useUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &useUserRepository{db}
}
func (ur *useUserRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(user).Error; nil != err {
		return err
	}
	return nil
}
func (ur *useUserRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; nil != err {
		return err
	}
	return nil
}
