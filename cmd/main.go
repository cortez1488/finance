package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"myFinanceTask/internal/core/auth"
	"myFinanceTask/internal/db/postgres/auth"
	"myFinanceTask/internal/handler/rest"
)

func main() {
	db := initPostgresDB()

	aRepo := psqlAuth.NewAuthStorage(db)
	aService := auth.NewAuthService(aRepo)
	handler := rest.NewHandler(aService)

	server := handler.InitRoutes()
	server.Run()
}

func initPostgresDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres password=qwerty dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
