package postgres

import (
	"ApiTest/internal/models"
	"fmt"
)

type UserDatabase struct {
	database *Database
}

func (u *UserDatabase) Create(user *models.User) (int64, error) {

	sqlStatemant := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	var id int64
	if err := u.database.db.Ping(); err != nil {
		fmt.Println("Не удалось подключиться к базе данных:", err)
		return -1, err
	}
	err := u.database.db.QueryRow(sqlStatemant, user.Name).Scan(&id)

	if err != nil {
		fmt.Println("Ошибка", err)
		return -1, err
	}
	return id, nil
}
