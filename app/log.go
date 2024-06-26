package app

import (
	"github.com/rs/zerolog"
	"os"
)

func (a *App) newLogger() {

	output := zerolog.ConsoleWriter{Out: os.Stderr}
	lg := zerolog.New(output).With().Timestamp().Logger()

	if a.Config.App.Environment == EnvironmentDevelopment {
		lg = lg.Level(zerolog.DebugLevel)
	} else {
		lg = lg.Level(zerolog.InfoLevel)
	}

	a.log = &lg
}

// Logger returns a configured instance of zerolog.Logger
func (a *App) Logger() *zerolog.Logger {
	return a.log
}
