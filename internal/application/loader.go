package application

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Database initialization
func LoadDB(driver, url string) (*gorm.DB, error) {
	return gorm.Open(driver, url)
}

// Load .env file from root directory
func LoadEnv() error {
	return godotenv.Load()
}

// Logrus logger initialization
func InitLogger(level string) *logrus.Logger {
	logger := logrus.New()

	loglevel, err := logrus.ParseLevel(level)
	if err != nil {
		panic("Error setting logger level")
	}

	logger.SetLevel(loglevel)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	return logger
}
