package database

import "ApiTest/internal/models"

type UserRepository interface {
	Create(user *models.User) (int64, error)
}

type WorkRepository interface {
	Create(work *models.Work) (int64, error)
}
