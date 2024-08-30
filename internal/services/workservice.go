package services

import (
	"ApiTest/internal/database"
	"ApiTest/internal/models"
)

type WorkService interface {
	CreateWork(user *models.Work) (int64, error)
}

type workServiceImpl struct {
	db database.Repository
}

func NewWorkService(db database.Repository) WorkService {
	return &workServiceImpl{db: db}
}

func (w *workServiceImpl) CreateWork(work *models.Work) (int64, error) {
	id, err := w.db.Work().Create(work)
	return id, err
}
