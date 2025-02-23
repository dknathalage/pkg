package command

import (
	"bytes"
	"flag"
	"io"
	"log/slog"
	"os"
	"strings"
	"testing"
)

// captureOutput captures and returns CLI output for testing.
func captureOutput(f func() error) (string, error) {
	r, w, _ := os.Pipe()
	os.Stdout = w // Redirect stdout

	err := f() // Execute function and capture error

	w.Close()
	os.Stdout = os.Stderr // Reset stdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r) // Read captured output
	return buf.String(), err
}

// TestNewCommandSet ensures the command set is initialized correctly.
func TestNewCommandSet(t *testing.T) {
	log := slog.Default()
	cs := NewCommandSet("app_name", log)
	if cs == nil {
		t.Fatal("Expected NewCommandSet() to return a non-nil value")
	}
	if len(cs.Commands) != 0 {
		t.Errorf("Expected an empty command set, got %d commands", len(cs.Commands))
	}
}

// TestRegisterCommand verifies command registration.
func TestRegisterCommand(t *testing.T) {
	log := slog.Default()
	cs := NewCommandSet("app_name", log)
	cs.RegisterCommand("test", "A test command")

	if len(cs.Commands) != 1 {
		t.Errorf("Expected 1 command, got %d", len(cs.Commands))
	}

	if _, exists := cs.Commands["test"]; !exists {
		t.Error("Command 'test' not found in command set")
	}
}

// TestRegisterSubcommand ensures subcommands are correctly registered.
func TestRegisterSubcommand(t *testing.T) {
	log := slog.Default()
	cs := NewCommandSet("app_name", log)
	cs.RegisterCommand("test", "A test command")
	cs.RegisterSubcommand("test", "sub", "A test subcommand", nil, func(args map[string]string) {})

	cmd, exists := cs.Commands["test"]
	if !exists {
		t.Fatal("Command 'test' not registered")
	}

	if len(cmd.Subcommands) != 1 {
		t.Errorf("Expected 1 subcommand, got %d", len(cmd.Subcommands))
	}

	if _, exists := cmd.Subcommands["sub"]; !exists {
		t.Error("Subcommand 'sub' not found in command 'test'")
	}
}

// TestRunWithInvalidCommand ensures Run handles unknown commands correctly.
func TestRunWithInvalidCommand(t *testing.T) {
	log := slog.Default()
	cs := NewCommandSet("app_name", log)
	os.Args = []string{"app_name", "invalid", "sub"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // Reset flag parsing

	output, err := captureOutput(cs.Run)

	if err == nil || !strings.Contains(err.Error(), "command 'invalid' not found") {
		t.Errorf("Expected error for unknown command, got: %v", err)
	}

	if !strings.Contains(output, "Error: Command 'invalid' not found") {
		t.Errorf("Expected error message in output, got: %s", output)
	}
}

// TestRunWithInvalidSubcommand ensures Run handles unknown subcommands correctly.
func TestRunWithInvalidSubcommand(t *testing.T) {
	log := slog.Default()
	cs := NewCommandSet("app_name", log)
	cs.RegisterCommand("test", "A test command")
	os.Args = []string{"app_name", "test", "invalid"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // Reset flag parsing

	output, err := captureOutput(cs.Run)

	if err == nil || !strings.Contains(err.Error(), "subcommand 'invalid' not found") {
		t.Errorf("Expected error for unknown subcommand, got: %v", err)
	}

	if !strings.Contains(output, "Error: Subcommand 'invalid' not found") {
		t.Errorf("Expected error message in output, got: %s", output)
	}
}

// TestRunValidSubcommand ensures a registered subcommand executes properly.
func TestRunValidSubcommand(t *testing.T) {
	log := slog.Default()
	cs := NewCommandSet("app_name", log)
	cs.RegisterCommand("test", "A test command")

	var executed bool
	cs.RegisterSubcommand("test", "run", "A test subcommand", nil, func(args map[string]string) {
		executed = true
	})

	os.Args = []string{"app_name", "test", "run"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // Reset flag parsing
	_, err := captureOutput(cs.Run)                                  // Capture CLI output to prevent pollution

	if err != nil {
		t.Errorf("Unexpected error during execution: %v", err)
	}

	if !executed {
		t.Error("Expected subcommand 'run' to execute but it did not")
	}
}

// TestFlagsParsing ensures flags are properly parsed and passed.
func TestFlagsParsing(t *testing.T) {
	log := slog.Default()
	cs := NewCommandSet("app_name", log)
	cs.RegisterCommand("test", "A test command")

	// Initialize receivedArgs properly
	receivedArgs := make(map[string]string)

	// Register a subcommand that updates receivedArgs
	cs.RegisterSubcommand("test", "run", "A test subcommand",
		map[string]string{"key": "default"}, func(args map[string]string) {
			for k, v := range args {
				receivedArgs[k] = v // Correctly store parsed arguments
			}
		})

	// Reset flag parsing for testing
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Manually register flags for test execution
	keyFlag := flag.String("key", "default", "Test flag for key")
	flag.Parse()

	// Simulate CLI arguments
	os.Args = []string{"app_name", "test", "run", "--key", "value"}
	_, err := captureOutput(cs.Run) // Prevent CLI output in test

	// Validate no unexpected errors
	if err != nil {
		t.Errorf("Unexpected error during flag parsing: %v", err)
	}

	// Validate correct flag parsing
	if *keyFlag != "value" {
		t.Errorf("Expected flag key to be 'value', got '%s'", *keyFlag)
	}
}
