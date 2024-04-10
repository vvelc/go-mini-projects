package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func New(isDebug bool) *zerolog.Logger {
	runLogFile, _ := os.OpenFile(
        "structured-api.log",
        os.O_APPEND|os.O_CREATE|os.O_WRONLY,
        0664,
    )
	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)

	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.TraceLevel
	}

	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(multi).With().Timestamp().Logger()

	return &logger
}
