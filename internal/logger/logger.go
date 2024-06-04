package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var (
	Log zerolog.Logger
)

func Info() *zerolog.Event {
	return Log.Info()
}

func Error() *zerolog.Event {
	return Log.Error()
}

func Fatal() *zerolog.Event {
	return Log.Fatal()
}

func InitiateLoggerToFile() func() {
	file, err := os.OpenFile(
		"splitflap.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		panic(err)
	}

	Log = zerolog.New(file).With().Timestamp().Logger()

	return func() { file.Close() }
}
