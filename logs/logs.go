package logs

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	fileLogger zerolog.Logger
	logger     zerolog.Logger
	colorGreen string = "\033[32m"
	colorReset string = "\033[0m"
	colorRed   string = "\033[31m"
)

func LogInfo(msg string) {
	logger.Info().Msg("ℹ️  " + string(colorGreen) + msg + string(colorReset))
	file, fileLogger := getFileLogger()
	fileLogger.Info().Msg(msg)
	defer file.Close()
}

func LogError(err error) {
	logger.Err(err).Msg("🚫  " + string(colorGreen) + "error encountered" + string(colorReset) + "\n\n")
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

		_, err = os.Create("logs/app.log")
		if err != nil {
			panic(err)
		}

		file, err = os.OpenFile(
			"logs/app.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		if err != nil {
			panic(err)
		}
	}

	fileLogger = zerolog.New(file).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Logger()
	return file, fileLogger
}

func init() {
	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		With().
		Timestamp().
		Logger()
}
