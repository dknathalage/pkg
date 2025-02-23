package app

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/dknathalage/pkg/cfg"
	"github.com/dknathalage/pkg/command"
)

// CliApp represents the root CLI application.
type CliApp[T any] struct {
	Name     string
	Logger   *slog.Logger
	Commands *command.CommandSet
	Config   T
}

const LevelCritical = slog.Level(12)

// NewCliApp initializes a new CLI application with logging and a command set.
func NewCliApp[T any](AppName string, Config T) *CliApp[T] {

	// creating a logging handler
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
				level := a.Value.Any().(slog.Level)
				if level == LevelCritical {
					a.Value = slog.StringValue("CRITICAL")
				}
			}
			return a
		},
	})

	cliapp := &CliApp[T]{
		Name:   AppName,
		Logger: slog.New(h),
		Config: Config,
	}

	cliapp.Commands = command.NewCommandSet(AppName, cliapp.Logger)

	return cliapp
}

// Run executes the CLI application.
func (app *CliApp[T]) Run() {
	if err := cfg.LoadEnvWithPrefix(strings.ToUpper(app.Name), app.Config); err != nil {
		app.Logger.Error(fmt.Sprintf("Error loading environment variables: %v", err))
	}
	app.Commands.Run()
}
