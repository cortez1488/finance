package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
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

func init() {
	initConfig()
}

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

	go func() {
		for {
			handler.RefreshPrices()
			refreshTime, err := time.ParseDuration(viper.Get("business.pricesRefillTime").(string))
			if err != nil {
				log.Fatalln("parsing refresh time error")
			}
			time.Sleep(refreshTime)
		}
	}()

	server.Run()
}

func initConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	viper.SetDefault("db.redis.password", "")
	viper.SetDefault("db.redis.db", 0)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("set config error")
	}
}

func initPostgresDB() *sqlx.DB {
	db, err := sqlx.Connect(viper.Get("db.postgres.drivername").(string), fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		viper.Get("db.postgres.username"), viper.Get("db.postgres.password"),
		viper.Get("db.postgres.dbname"), viper.Get("db.postgres.sslmode")))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func initRedisDB() *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     viper.Get("db.redis.address").(string),
		Password: viper.Get("db.redis.password").(string), // no password set
		DB:       viper.Get("db.redis.db").(int),          // use default DB
	})
	return db
}
