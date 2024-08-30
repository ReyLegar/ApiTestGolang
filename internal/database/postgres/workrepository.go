package postgres

import (
	"ApiTest/internal/models"
	"fmt"
)

type WorkDatabase struct {
	database *Database
}

func (u *WorkDatabase) Create(work *models.Work) (int64, error) {

	sqlStatemant := `INSERT INTO works (title) VALUES ($1) RETURNING id`
	var id int64
	if err := u.database.db.Ping(); err != nil {
		fmt.Println("Не удалось подключиться к базе данных:", err)
		return -1, err
	}
	err := u.database.db.QueryRow(sqlStatemant, work.Title).Scan(&id)

	if err != nil {
		fmt.Println("Ошибка", err)
		return -1, err
	}
	return id, nil
}
