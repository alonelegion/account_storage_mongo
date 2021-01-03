package recoverywriter

import (
	"errors"
	"github.com/sirupsen/logrus"
	"strings"
)

type GinRecoverWriter struct {
	logger *logrus.Logger
}

func NewGinRecoverWriter(logger *logrus.Logger) *GinRecoverWriter {
	return &GinRecoverWriter{logger: logger}
}

func (w *GinRecoverWriter) Write(p []byte) (n int, err error) {
	w.logger.WithError(errors.New(strings.TrimSpace(string(p)))).Error("Gin Recovered")

	return 0, nil
}
