package application

import (
	"context"
	"errors"
	"github.com/alonelegion/account_storage_mongo/internal/cache"
	"github.com/alonelegion/account_storage_mongo/internal/queries"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/now"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Application struct {
	db           *gorm.DB
	logger       *logrus.Logger
	envMode      EnvironmentMode
	shutdowners  []Shutdown
	Dependencies *Dependencies
}

type Dependencies struct {
	BaseQuery queries.Query
	Storage   cache.Storager
}

var (
	// Current Application
	instance *Application
	once     sync.Once
)

// Application errors
var (
	ErrInvalidEnvMode = errors.New("invalid environment mode")
)

// Initializing a Application
func Init() *Application {
	once.Do(func() {
		envMode, errEnv := receiveEnvironmentMode()
		if errEnv != nil {
			logrus.Warn(errEnv)
		}

		instance = &Application{
			envMode:      envMode,
			Dependencies: &Dependencies{},
		}

		// set default week start day to monday
		now.WeekStartDay = time.Monday
	})
	return instance
}

func (a *Application) AddShutdown(s Shutdown) {
	a.shutdowners = append(a.shutdowners, s)
}

func (a *Application) Shutdown(ctx context.Context) {
	wg := &sync.WaitGroup{}

	for _, s := range a.shutdowners {
		wg.Add(1)

		go func(s Shutdown) {
			if err := s(ctx); err != nil {
				a.logger.Error(err)
			}

			wg.Done()
		}(s)
	}

	wg.Wait()

	defer a.CloseConnections()
}

func (a *Application) CloseConnections() {
	if a.db != nil {
		if err := a.db.Close(); err != nil {
			a.logger.Error(err)
		}
	}
}

func Get() *Application {
	return instance
}

func (a *Application) DB() *gorm.DB {
	return a.db
}

func (a *Application) SetDB(db *gorm.DB) {
	a.db = db
}

func (a *Application) Logger() *logrus.Logger {
	return a.logger
}

func (a *Application) SetLogger(logger *logrus.Logger) {
	a.logger = logger
}

func (a *Application) EnvMode() EnvironmentMode {
	return a.envMode
}

func (a *Application) Deps() *Dependencies {
	return a.Dependencies
}

func Logger() *logrus.Logger {
	return Get().Logger()
}

func DB() *gorm.DB {
	return Get().DB()
}

func EnvMode() EnvironmentMode {
	return Get().EnvMode()
}

func Deps() *Dependencies {
	return Get().Deps()
}

func Storage() cache.Storager {
	return Deps().Storage
}
