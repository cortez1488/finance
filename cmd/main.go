package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	"os"
	"strings"
	"time"
)

func init() {
	initConfig()
}

func main() {
	rdb := initRedisDB()
	db := initPostgresDB()
	migrateDB(db)

	log.Println("Init application logic")
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
	viper.SetDefault("db.redis.host", "localhost")

	viper.SetDefault("db.postgres.port", "5432")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("set config error")
	}
}

func initPostgresDB() *sqlx.DB {
	log.Println("sql connect string:", getPostgresDBConnectString())
	var db *sqlx.DB
	var err error
	var errConnectionRefusedCounter int

	for {
		db, err = sqlx.Connect(viper.Get("db.postgres.drivername").(string), getPostgresDBConnectString())

		if err != nil {
			if strings.Contains(err.Error(), "connect: connection refused") {
				errConnectionRefusedCounter++
				time.Sleep(time.Millisecond * 500)
				log.Println(errConnectionRefusedCounter+1, "attempt to connect to database")
			}

			if errConnectionRefusedCounter >= 5 {
				log.Fatal("PostgresDB initialization " + err.Error())
			}
		} else {
			break
		}
	}
	return db
}

func getPostgresDBConnectString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), viper.Get("db.postgres.port"),
		viper.Get("db.postgres.username"), viper.Get("db.postgres.password"),
		viper.Get("db.postgres.dbname"), viper.Get("db.postgres.sslmode"))
}

func migrateDB(db *sqlx.DB) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:./migrations/",
		viper.GetString("dbname"), driver)
	if err != nil {
		log.Fatalln("Error with database migration creating:", err)
	}

	err = m.Up()
	if err != nil {
		if !strings.Contains(err.Error(), "no change") {
			log.Fatalln("Error with database migration:", err)
		}
	}
}

func initRedisDB() *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CACHE_HOST"),
		Password: viper.Get("db.redis.password").(string), // no password set
		DB:       viper.Get("db.redis.db").(int),          // use default DB
	})
	return db
}
