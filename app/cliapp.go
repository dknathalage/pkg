package app

import (
	"github.com/dknathalage/pkg/command"
	"github.com/dknathalage/pkg/log"
)

// CliApp represents the root CLI application.
type CliApp struct {
	Log      *log.Logger
	Commands *command.CommandSet
}

// NewCliApp initializes a new CLI application with logging and a command set.
func NewCliApp() *CliApp {
	return &CliApp{
		Log:      log.NewLogger(),
		Commands: command.NewCommandSet(),
	}
}

// Run executes the CLI application.
func (app *CliApp) Run() {
	app.Commands.Run()
}
