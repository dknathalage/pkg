package app

import (
	"fmt"
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
	if Config == nil || reflect.TypeOf(Config).Kind() != reflect.Ptr || reflect.TypeOf(Config).Elem().Kind() != reflect.Struct {
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
	if err := LoadEnvWithPrefix(strings.ToUpper(app.Name), app.Config); err != nil {
		app.Logger.Error(fmt.Sprintf("Error loading environment variables: %v", err))
	}
	app.Commands.Run()
}
