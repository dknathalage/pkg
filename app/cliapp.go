package app

import (
	"reflect"
	"strings"

	"github.com/dknathalage/pkg/command"
	"github.com/dknathalage/pkg/log"
)

// CliApp represents the root CLI application.
type CliApp struct {
	Name     string
	Logger   *log.Logger
	Commands *command.CommandSet
	Config   interface{}
}

// NewCliApp initializes a new CLI application with logging and a command set.
func NewCliApp(AppName string, Config interface{}) *CliApp {

	// Ensure Config is a pointer to a struct
	if reflect.TypeOf(Config).Kind() != reflect.Ptr {
		panic("Config must be a pointer to a struct")
	}

	cliapp := &CliApp{
		Name:   AppName,
		Logger: log.NewLogger(),
		Config: Config,
	}

	cliapp.Commands = command.NewCommandSet(AppName, cliapp.Logger)

	return cliapp
}

// Run executes the CLI application.
func (app *CliApp) Run() {
	LoadEnvWithPrefix(strings.ToUpper(app.Name), app.Config)
	app.Commands.Run()
}
