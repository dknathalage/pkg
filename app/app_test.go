package app

import (
	"testing"
)

func TestNewAppLogger(t *testing.T) {
	app := NewApp()
	if app.Logger == nil {
		t.Fatal("Expected logger to be initialized, got nil")
	}
}

func TestNewAppCli(t *testing.T) {
	app := NewApp()
	if app.Cli == nil {
		t.Fatal("Expected cli to be initialized, got nil")
	}
}
