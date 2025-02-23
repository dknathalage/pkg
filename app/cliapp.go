package app

import (
	"fmt"
	"log/slog"
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

// NewCliApp initializes a new CLI application with logging and a command set.
func NewCliApp[T any](AppName string, Config T) *CliApp[T] {
	cliapp := &CliApp[T]{
		Name:   AppName,
		Logger: slog.Default(),
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
