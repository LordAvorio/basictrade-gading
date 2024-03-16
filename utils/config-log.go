package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitZeroLog() {

	today := time.Now().Format("02-01-2006")

	logFolder := "logs"

	if err := os.MkdirAll(logFolder, 0755); err != nil {
		panic("Failed to create log folder")
	}

	nameLog := fmt.Sprintf("logfile_basictrade_%s.log", today)
	logFilePath := filepath.Join(logFolder, nameLog)

	file, err := os.OpenFile(
		logFilePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)

	if err != nil {
		panic(err)
	}

	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = zerolog.New(file).With().Timestamp().Caller().Logger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

}
