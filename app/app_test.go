package app

import (
	"testing"
)

func TestNewApp(t *testing.T) {
	app := NewApp()
	if app.Logger == nil {
		t.Fatal("Expected logger to be initialized, got nil")
	}
}
