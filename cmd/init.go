package main

import (
	"fmt"
	"github.com/alonelegion/account_storage_mongo/internal/application"
	"github.com/alonelegion/account_storage_mongo/internal/queries"
	"github.com/alonelegion/account_storage_mongo/internal/router"
	"github.com/alonelegion/account_storage_mongo/internal/runners"
	"net/http"
	"os"
)

func InitLogger(app *application.Application) {
	loggerLevel := os.Getenv("LOGGER_LEVEL")
	if loggerLevel == "" {
		loggerLevel = "debug"
	}

	logger := application.InitLogger(loggerLevel)

	app.SetLogger(logger)
}

func InitDB(app *application.Application) {
	dbURL := os.Getenv("DATABASE_DRIVER")
	if prodURL := os.Getenv("PROD_DATABASE_URL"); prodURL != "" {
		dbURL = prodURL
	}

	db, err := application.LoadDB(
		os.Getenv("DATABASE_DRIVER"), dbURL+"?sslmode=disable",
	)
	if err != nil {
		app.Logger().Panic(err)
	}

	if os.Getenv("DB_DEBUG") == "true" {
		db = db.Debug()
	}

	db.SetLogger(app.Logger())

	app.SetDB(db)
	app.Dependencies.BaseQuery = queries.InitQuery(db)
}

func InitWebServer(app *application.Application) {
	engine := router.NewRouter(app.Logger())

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler: engine,
	}

	app.AddShutdown(server.Shutdown)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			application.Logger().Panic(err)
		}
	}()
}

func InitRoutines(app *application.Application) {
	erh := runners.ExternalRoutine()
	app.AddShutdown(erh.Shutdown)
}
