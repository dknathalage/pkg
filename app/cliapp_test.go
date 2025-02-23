package app

import (
	"os"
	"testing"

	"github.com/dknathalage/pkg/log"
)

type TestConfig struct {
	Test string
}

func TestNewCliApp(t *testing.T) {
	os.Setenv("APP_NAME_TEST", "hello")

	app := NewCliApp("app_name", &TestConfig{})

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
