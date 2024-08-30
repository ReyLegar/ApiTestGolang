package database

type Repository interface {
	User() UserRepository
	Work() WorkRepository
}
