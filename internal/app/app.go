package app

import (
	"ApiTest/config"
	"ApiTest/internal/database/postgres"
	"ApiTest/internal/services"
	restTransport "ApiTest/internal/transport/rest"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Run(cfg *config.Config) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DB.Host, cfg.DB.DBPort, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	db, err := sql.Open("postgres", connStr)
	fmt.Println(connStr)
	if err != nil {
		panic(err)
	}

	postgresDB := postgres.New(db)

	userService := services.NewUserService(postgresDB)
	workService := services.NewWorkService(postgresDB)

	router := restTransport.NewRouter(userService, workService)

	server := restTransport.NewServer(cfg.Port, router)
	server.Start()
}
