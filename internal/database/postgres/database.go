package postgres

import (
	"ApiTest/internal/database"
	"database/sql"
)

type Database struct {
	db *sql.DB

	userDatabase *UserDatabase
	workDatabase *WorkDatabase
}

func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

func (db *Database) User() database.UserRepository {
	if db.userDatabase != nil {
		return db.userDatabase
	}

	db.userDatabase = &UserDatabase{
		database: db,
	}

	return db.userDatabase
}

func (db *Database) Work() database.WorkRepository {
	if db.workDatabase != nil {
		return db.workDatabase
	}

	db.workDatabase = &WorkDatabase{
		database: db,
	}

	return db.workDatabase
}
