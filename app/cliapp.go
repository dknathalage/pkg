package app

import (
	"github.com/dknathalage/pkg/command"
	"github.com/dknathalage/pkg/log"
)

// CliApp represents the root CLI application.
type CliApp struct {
	Logger   *log.Logger
	Commands *command.CommandSet
}

// NewCliApp initializes a new CLI application with logging and a command set.
func NewCliApp(AppName string) *CliApp {
	cliapp := &CliApp{
		Logger: log.NewLogger(),
	}

	cliapp.Commands = command.NewCommandSet(AppName, cliapp.Logger)

	return cliapp
}

// Run executes the CLI application.
func (app *CliApp) Run() {
	app.Commands.Run()
}
