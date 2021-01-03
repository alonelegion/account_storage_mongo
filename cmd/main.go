package main

import (
	"context"
	"github.com/alonelegion/account_storage_mongo/internal/application"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	appShutdownDuration = 20 * time.Second
)

func main() {
	// load os.env settings
	if err := application.LoadEnv(); err != nil {
		logrus.Warning(".env file not find")
	}

	// init main application singleton instance
	app := application.Init()

	InitLogger(app)
	InitDB(app)

	// if something goes wrong, catch it
	defer func() {
		if r := recover(); r != nil {
			app.Logger().
				WithField("RecoverMessage", r).
				Fatal("Application Crashed")
		}
	}()

	InitRoutines(app)
	InitWebServer(app)

	// shutdown
	<-GracefulShutdown()
	app.Logger().Debug("SIGTERM signal received. Shutdown...")

	ctx, forceCancel := context.WithTimeout(context.Background(), appShutdownDuration)
	defer forceCancel()

	app.Shutdown(ctx)
}

func GracefulShutdown() chan os.Signal {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return done
}
