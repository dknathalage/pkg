package log

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestNewJsonLogger(t *testing.T) {
	r, w, _ := os.Pipe()

	oldStderr := os.Stderr

	os.Stderr = w

	logger := NewJsonLogger()
	logger.Info("Test log message")

	w.Close()
	os.Stderr = oldStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	logOutput := buf.String()

	if !strings.Contains(logOutput, "Test log message") {
		t.Errorf("Expected 'message' key in log output, got: %s", logOutput)
	}

	if !strings.Contains(logOutput, "INFO") {
		t.Errorf("Expected 'severity' key in log output, got: %s", logOutput)
	}
}
