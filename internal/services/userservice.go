package services

import (
	"ApiTest/internal/database"
	"ApiTest/internal/models"
)

type UserService interface {
	CreateUser(user *models.User) (int64, error)
}

type userServiceImpl struct {
	db database.Repository
}

func NewUserService(db database.Repository) UserService {
	return &userServiceImpl{db: db}
}

func (u *userServiceImpl) CreateUser(user *models.User) (int64, error) {
	id, err := u.db.User().Create(user)
	return id, err
}
