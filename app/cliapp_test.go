package app

import (
	"testing"

	"github.com/dknathalage/pkg/log"
)

func TestNewCliApp(t *testing.T) {
	app := NewCliApp("app_name")

	if app == nil {
		t.Fatal("Expected CliApp instance, got nil")
	}

	if app.Logger == nil {
		t.Fatal("Expected logger to be initialized, got nil")
	}

	_, ok := interface{}(app.Logger).(*log.Logger)
	if !ok {
		t.Fatal("Expected app.Log to be of type *log.Logger")
	}
}
