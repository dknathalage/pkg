package app

import (
	"testing"

	"github.com/dknathalage/pkg/log"
)

// TestNewCliApp ensures that NewCliApp correctly initializes a CliApp instance.
func TestNewCliApp(t *testing.T) {
	app := NewCliApp()

	if app == nil {
		t.Fatal("Expected CliApp instance, got nil")
	}

	if app.Log == nil {
		t.Fatal("Expected logger to be initialized, got nil")
	}

	_, ok := interface{}(app.Log).(*log.Logger)
	if !ok {
		t.Fatal("Expected app.Log to be of type *log.Logger")
	}
}
