package repository

import (
	"gorm.io/gorm"
	"user_service/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser() ([]*models.User, error)
	GetUserByID(id int) (*models.User, error)
	DeleteUser(user *models.User) error
}

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) CreateUser(user *models.User) error {
	return u.DB.Create(&user).Error
}

func (u *UserRepo) GetUser() ([]*models.User, error) {
	var res []*models.User
	err := u.DB.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *UserRepo) GetUserByID(id int) (*models.User, error) {
	var res *models.User
	err := u.DB.First(&res, id).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *UserRepo) DeleteUser(user *models.User) error {
	return u.DB.Delete(&user).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepo{DB: db}
}
