package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(level string, prettyConsole bool) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		zerolog.SetGlobalLevel(lvl)
	}

	log.Logger = log.With().Caller().Logger()

	if prettyConsole {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
	}

	log.Info().Msg("Logger initialized")
}
