package application

import "os"

type EnvironmentMode string

const (
	Unknown     EnvironmentMode = "unknown"
	Development EnvironmentMode = "dev"
	Production  EnvironmentMode = "prod"
)

// Receive and validate environment mode from .env file
func receiveEnvironmentMode() (EnvironmentMode, error) {
	envMode := EnvironmentMode(os.Getenv("APP_ENV"))

	switch envMode {
	case Development, Production:
		return envMode, nil
	}

	return Unknown, ErrInvalidEnvMode
}
