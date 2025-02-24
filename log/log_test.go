package log

import (
	"bytes"
	"log/slog"
	"os"
	"strings"
	"testing"
)

func captureLogOutput(f func(logger *slog.Logger)) string {
	// Create a pipe
	r, w, _ := os.Pipe()
	oldStderr := os.Stderr
	os.Stderr = w

	// Run the logger function
	logger := NewJsonLogHandler()
	f(logger)

	// Close writer to flush, restore stderr
	w.Close()
	os.Stderr = oldStderr

	// Read log output
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	return buf.String()
}

func TestNewJsonLogHandler(t *testing.T) {
	// Create a pipe
	r, w, _ := os.Pipe()

	// Save the original stderr
	oldStderr := os.Stderr

	// Redirect stderr to our pipe
	os.Stderr = w

	// Call the logger
	logger := NewJsonLogHandler()
	logger.Info("Test log message")

	// Close writer to flush data and restore stderr
	w.Close()
	os.Stderr = oldStderr

	// Read output from pipe
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	// Log and check the output
	logOutput := buf.String()

	if !strings.Contains(logOutput, "Test log message") {
		t.Errorf("Expected 'message' key in log output, got: %s", logOutput)
	}

	if !strings.Contains(logOutput, "ERROR") {
		t.Errorf("Expected 'severity' key in log output, got: %s", logOutput)
	}
}
