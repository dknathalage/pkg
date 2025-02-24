package app

import (
	"testing"

	"github.com/urfave/cli/v3"
)

func TestNewAppLogger(t *testing.T) {
	app := NewApp(&cli.Command{})
	if app.Logger == nil {
		t.Fatal("Expected logger to be initialized, got nil")
	}
}

func TestNewAppCli(t *testing.T) {
	app := NewApp(&cli.Command{})
	if app.Cli == nil {
		t.Fatal("Expected cli to be initialized, got nil")
	}
}
