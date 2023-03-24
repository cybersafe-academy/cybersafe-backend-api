package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func Config(panicPrefix, appName, appVersion string, production bool) *zerolog.Logger {
	// Logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	if os.Getenv("ZEROLOG_PRETTY_PRINT") != "" {
		logger = logger.Output(zerolog.NewConsoleWriter())
	}

	return &logger
}
