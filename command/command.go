package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/dknathalage/pkg/log"
)

// Command represents a CLI command with subcommands.
type Command struct {
	Name        string
	Description string
	Subcommands map[string]*Subcommand
}

// Subcommand represents a CLI subcommand with flags and execution logic.
type Subcommand struct {
	Name        string
	Description string
	Flags       map[string]*string
	Execute     func(args map[string]string)
}

// CommandSet holds all registered commands.
type CommandSet struct {
	Name     string
	logger   log.Logger
	Commands map[string]*Command
}

// NewCommandSet initializes a new command set.
func NewCommandSet(Name string, Logger *log.Logger) *CommandSet {
	return &CommandSet{
		Name:     Name,
		logger:   *Logger,
		Commands: make(map[string]*Command),
	}
}

// RegisterCommand adds a top-level command.
func (cs *CommandSet) RegisterCommand(name, description string) {
	cs.Commands[name] = &Command{
		Name:        name,
		Description: description,
		Subcommands: make(map[string]*Subcommand),
	}
}

// RegisterSubcommand adds a subcommand under a command.
func (cs *CommandSet) RegisterSubcommand(commandName, subcommandName, description string, flags map[string]string, execute func(args map[string]string)) {
	cmd, exists := cs.Commands[commandName]
	if !exists {
		fmt.Printf("Error: Command '%s' not found\n", commandName)
		return
	}

	subcommand := &Subcommand{
		Name:        subcommandName,
		Description: description,
		Flags:       make(map[string]*string),
		Execute:     execute,
	}

	// Register flags for the subcommand
	for flagName, defaultValue := range flags {
		subcommand.Flags[flagName] = flag.String(flagName, defaultValue, fmt.Sprintf("Flag for %s", flagName))
	}

	cmd.Subcommands[subcommandName] = subcommand
}

// Run executes the CLI by parsing arguments and dispatching the correct subcommand.
func (cs *CommandSet) Run() error {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <command> <subcommand> [--flags]", cs.Name)
		cs.PrintUsage()
		return fmt.Errorf("missing command or subcommand")
	}

	commandName := os.Args[1]
	subcommandName := os.Args[2]

	command, exists := cs.Commands[commandName]
	if !exists {
		fmt.Printf("Error: Command '%s' not found\n", commandName)
		cs.PrintUsage()
		return fmt.Errorf("command '%s' not found", commandName)
	}

	subcommand, exists := command.Subcommands[subcommandName]
	if !exists {
		fmt.Printf("Error: Subcommand '%s' not found in command '%s'\n", subcommandName, commandName)
		cs.PrintUsage()
		return fmt.Errorf("subcommand '%s' not found", subcommandName)
	}

	// Parse flags
	flag.CommandLine.Parse(os.Args[3:])

	// Collect flag values into a map
	args := make(map[string]string)
	for key, value := range subcommand.Flags {
		args[key] = *value
	}

	subcommand.Execute(args)
	return nil
}

// PrintUsage displays available commands and subcommands.
func (cs *CommandSet) PrintUsage() {
	fmt.Println("Available commands:")
	for cmdName, cmd := range cs.Commands {
		fmt.Printf("  %s: %s\n", cmdName, cmd.Description)
		for subName, sub := range cmd.Subcommands {
			fmt.Printf("    - %s: %s\n", subName, sub.Description)
		}
	}
}
