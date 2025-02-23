package log

import (
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"
)

func captureLogOutput(f func(l *Logger)) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	l := NewLogger()
	f(l)

	w.Close()
	os.Stdout = stdout

	var sb strings.Builder
	_, _ = io.Copy(&sb, r)
	return sb.String()
}

func TestLoggingLevels(t *testing.T) {
	tests := []struct {
		level   string
		message string
		logFunc func(*Logger, string)
	}{
		{INFO, "Info message", func(l *Logger, msg string) { l.Info(msg) }},
		{DEBUG, "Debug message", func(l *Logger, msg string) { l.Debug(msg) }},
		{WARNING, "Warning message", func(l *Logger, msg string) { l.Warning(msg) }},
		{ERROR, "Error message", func(l *Logger, msg string) { l.Error(msg) }},
		{CRITICAL, "Critical message", func(l *Logger, msg string) { l.Critical(msg) }},
		{ALERT, "Alert message", func(l *Logger, msg string) { l.Alert(msg) }},
	}

	for _, tt := range tests {
		t.Run(tt.level, func(t *testing.T) {
			output := captureLogOutput(func(l *Logger) {
				tt.logFunc(l, tt.message)
			})

			var logEntry LogEntry
			if err := json.Unmarshal([]byte(output), &logEntry); err != nil {
				t.Fatalf("Failed to unmarshal log output: %v", err)
			}

			if logEntry.Severity != tt.level {
				t.Errorf("Expected severity %s, got %s", tt.level, logEntry.Severity)
			}

			if logEntry.Message != tt.message {
				t.Errorf("Expected message %s, got %s", tt.message, logEntry.Message)
			}

			if logEntry.Timestamp == "" {
				t.Error("Expected timestamp, but got empty")
			}
		})
	}
}

func TestLoggingWithContext(t *testing.T) {
	output := captureLogOutput(func(l *Logger) {
		l.Error("Database connection failed", map[string]interface{}{
			"database": "PostgreSQL",
			"host":     "127.0.0.1",
			"port":     5432,
		})
	})

	var logEntry LogEntry
	if err := json.Unmarshal([]byte(output), &logEntry); err != nil {
		t.Fatalf("Failed to unmarshal log output: %v", err)
	}

	if logEntry.Severity != ERROR {
		t.Errorf("Expected severity %s, got %s", ERROR, logEntry.Severity)
	}

	if logEntry.Message != "Database connection failed" {
		t.Errorf("Expected message 'Database connection failed', got '%s'", logEntry.Message)
	}

	if logEntry.Context["database"] != "PostgreSQL" {
		t.Errorf("Expected database 'PostgreSQL', got '%v'", logEntry.Context["database"])
	}

	if logEntry.Context["host"] != "127.0.0.1" {
		t.Errorf("Expected host '127.0.0.1', got '%v'", logEntry.Context["host"])
	}

	if logEntry.Context["port"] != float64(5432) {
		t.Errorf("Expected port 5432, got '%v'", logEntry.Context["port"])
	}
}

func TestLoggingCallerInfo(t *testing.T) {
	output := captureLogOutput(func(l *Logger) {
		l.Error("Testing caller info")
	})

	var logEntry LogEntry
	if err := json.Unmarshal([]byte(output), &logEntry); err != nil {
		t.Fatalf("Failed to unmarshal log output: %v", err)
	}

	if logEntry.Caller == "" {
		t.Error("Expected caller info, but got empty")
	}
}
