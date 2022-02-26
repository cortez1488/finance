package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"myFinanceTask/internal/core/admSymbol"
	"myFinanceTask/internal/core/auth"
	psqlAdmSymbol "myFinanceTask/internal/db/postgres/admSymbol"
	"myFinanceTask/internal/db/postgres/auth"
	"myFinanceTask/internal/handler/rest"
)

func main() {
	db := initPostgresDB()

	authRepo := psqlAuth.NewAuthStorage(db)
	authService := auth.NewAuthService(authRepo)

	admSymbolRepo := psqlAdmSymbol.NewAdmSymbolStorage(db)
	admSymbolService := admSymbol.NewAdmSymbolService(admSymbolRepo)

	handler := rest.NewHandler(authService, admSymbolService)

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
