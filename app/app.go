package app

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/urfave/cli/v3"
)

type App struct {
	Logger *slog.Logger
	Cli    *cli.Command
}

func (app *App) Run() {
	if err := app.Cli.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func NewApp(cliCommand *cli.Command) *App {
	return &App{
		Logger: slog.New(newJsonLogHandler()),
		Cli:    cliCommand,
	}
}

func newJsonLogHandler() *slog.JSONHandler {
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.MessageKey {
				a.Key = "message"
			} else if a.Key == slog.SourceKey {
				a.Key = "logging.googleapis.com/sourceLocation"
			} else if a.Key == slog.LevelKey {
				a.Key = "severity"
			}
			return a
		},
	})

	return h
}
