package cfg

import (
	"os"
	"reflect"
	"testing"
)

type Config struct {
	DatabaseURL string
	Port        int
	DebugMode   bool
}

func TestLoadEnvWithPrefix(t *testing.T) {
	// Set environment variables
	os.Setenv("APP_DATABASEURL", "postgres://user:pass@localhost:5432/db")
	os.Setenv("APP_PORT", "8080")
	os.Setenv("APP_DEBUGMODE", "true")

	expectedConfig := Config{
		DatabaseURL: "postgres://user:pass@localhost:5432/db",
		Port:        8080,
		DebugMode:   true,
	}

	config := Config{}
	LoadEnvWithPrefix("APP", &config)

	if !reflect.DeepEqual(config, expectedConfig) {
		t.Errorf("Expected %+v, got %+v", expectedConfig, config)
	}

	// Unset environment variables after test
	os.Unsetenv("APP_DATABASEURL")
	os.Unsetenv("APP_PORT")
	os.Unsetenv("APP_DEBUGMODE")
}

func TestLoadEnvWithPrefix_MissingValues(t *testing.T) {
	// Ensure no environment variables are set
	os.Unsetenv("APP_DATABASEURL")
	os.Unsetenv("APP_PORT")
	os.Unsetenv("APP_DEBUGMODE")

	expectedConfig := Config{} // Default values (empty string, 0, false)
	config := Config{}
	LoadEnvWithPrefix("APP", &config)

	if !reflect.DeepEqual(config, expectedConfig) {
		t.Errorf("Expected %+v, got %+v", expectedConfig, config)
	}
}

func TestLoadEnvWithPrefix_InvalidInt(t *testing.T) {
	// Set an invalid integer value
	os.Setenv("APP_PORT", "invalid")

	config := Config{}
	LoadEnvWithPrefix("APP", &config)

	if config.Port != 0 { // Default int value should be 0
		t.Errorf("Expected Port to be 0, got %d", config.Port)
	}

	os.Unsetenv("APP_PORT")
}

func TestLoadEnvWithPrefix_InvalidBool(t *testing.T) {
	// Set an invalid boolean value
	os.Setenv("APP_DEBUGMODE", "notaboolean")

	config := Config{}
	LoadEnvWithPrefix("APP", &config)

	if config.DebugMode { // Default bool value should be false
		t.Errorf("Expected DebugMode to be false, got true")
	}

	os.Unsetenv("APP_DEBUGMODE")
}
