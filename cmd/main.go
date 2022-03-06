package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"myFinanceTask/internal/core/admSymbol"
	"myFinanceTask/internal/core/auth"
	"myFinanceTask/internal/core/deal"
	"myFinanceTask/internal/core/price_refresh"
	"myFinanceTask/internal/core/user_account"
	psqlAdmSymbol "myFinanceTask/internal/db/admSymbol"
	"myFinanceTask/internal/db/auth"
	"myFinanceTask/internal/db/deal"
	price_refresh_storage "myFinanceTask/internal/db/price_refresh"
	"myFinanceTask/internal/db/userAccount"
	"myFinanceTask/internal/handler/rest"
	"time"
)

func main() {
	db := initPostgresDB()
	rdb := initRedisDB()

	authRepo := psqlAuth.NewAuthStorage(db)
	authService := auth.NewAuthService(authRepo)

	admSymbolRepo := psqlAdmSymbol.NewAdmSymbolStorage(db, rdb)
	admSymbolService := admSymbol.NewAdmSymbolService(admSymbolRepo)

	userAccountRepo := userAccount.NewUserAccountStorage(db)
	userAccountService := user_account.NewUserAccountService(userAccountRepo)

	dealRepo := dealStorage.NewDealStorage(db, rdb)
	dealService := deal.NewDealService(dealRepo)

	priceRefreshRepo := price_refresh_storage.NewPriceRefreshStorage(db, rdb)
	priceRefreshService := price_refresh.NewPriceRefreshService(priceRefreshRepo)

	handler := rest.NewHandler(authService, admSymbolService, userAccountService, dealService, priceRefreshService)

	server := handler.InitRoutes()

	rdb.Set(context.Background(), "test", "test", 0)
	log.Println("we did it")

	go func() {
		for {
			handler.RefreshPrices()
			time.Sleep(time.Second * 3)
		}
	}()

	server.Run()
}

func initPostgresDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres password=qwerty dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func initRedisDB() *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return db
}
