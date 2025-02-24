package log

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestNewJsonLogger(t *testing.T) {
	// Create a pipe
	r, w, _ := os.Pipe()

	// Save the original stderr
	oldStderr := os.Stderr

	// Redirect stderr to our pipe
	os.Stderr = w

	// Call the logger
	logger := NewJsonLogger()
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

	if !strings.Contains(logOutput, "INFO") {
		t.Errorf("Expected 'severity' key in log output, got: %s", logOutput)
	}
}
