package logs

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
)

var fileLogger zerolog.Logger
var logger zerolog.Logger

func LogInfo(msg string) {
	logger.Info().Msg(msg)
	file, fileLogger := getFileLogger()
	fileLogger.Info().Msg(msg)
	defer file.Close()
}

func LogError(err error) {
	logger.Err(err).Msg("error encountered")
	file, fileLogger := getFileLogger()
	fileLogger.Err(err).Msg("error encountered")
	defer file.Close()
}

func getFileLogger() (*os.File, zerolog.Logger) {

	file, err := os.OpenFile(
		"logs/app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		err := os.Mkdir("logs", os.ModePerm)
		if err != nil {
			panic(err)
		}

		logfile, err := os.Create("logs/app.log")
		if err != nil {
			panic(err)
		}
		defer logfile.Close()

		file, err = os.OpenFile(
			"logs/app.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		if err != nil {
			panic(err)
		}
	}

	defer file.Close()

	fileLogger = zerolog.New(file).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Logger()
	return file, fileLogger
}

func init() {
	buildInfo, _ := debug.ReadBuildInfo()
	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()
}
