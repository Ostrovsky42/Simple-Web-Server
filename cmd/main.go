package main

import (
	"SimplyWebServer/app"
	_ "SimplyWebServer/docs"
	"SimplyWebServer/config"
	"SimplyWebServer/infrastructure/handlers"
	"SimplyWebServer/infrastructure/storage"
	"SimplyWebServer/migrations"
	"context"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	bin "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title        Simply Web Server API
// @version      1.0
// @description  This is a sample web server for replenishment and transfer of funds.

// @host         localhost:8088
// @BasePath     /


func main() {
	e := echo.New()
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
	}

	config := config.NewConfig()

	db := ConnectDB(config)
	storage := storage.NewStorage(db, log)
	service := app.NewTransactionService(storage, log)
	handler := handlers.NewTransactionHandler(service, log)

	e.POST("/add_user", handlers.CreateUserProc(storage))
	e.POST("/add", handler.AddToBalance)
	e.POST("/transfer", handler.AddTransfer)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	migrateDB(config.DSN)

	go func() {
		if err := e.Start(":8088"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func ConnectDB(conf *config.Config) (db *sql.DB) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		logrus.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		logrus.Fatal(err)
	}
	return
}

func migrateDB(dsn string) {
	s := bin.Resource(migrations.AssetNames(), migrations.Asset)
	d, err := bin.WithInstance(s)
	if err != nil {
		logrus.Fatal(err)
	}
	m, err := migrate.NewWithSourceInstance("go-bindata", d, dsn)
	if err != nil {
		logrus.Fatal(err)
	}

	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			logrus.Warn(err)
		} else {
			logrus.Fatal(err)
		}
	}
}
