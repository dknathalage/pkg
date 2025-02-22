package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

// Log levels
const (
	DEFAULT   = "DEFAULT"
	DEBUG     = "DEBUG"
	INFO      = "INFO"
	NOTICE    = "NOTICE"
	WARNING   = "WARNING"
	ERROR     = "ERROR"
	CRITICAL  = "CRITICAL"
	ALERT     = "ALERT"
	EMERGENCY = "EMERGENCY"
)

// LogEntry represents a structured log message.
type LogEntry struct {
	Severity  string                 `json:"severity"`
	Message   string                 `json:"message"`
	Timestamp string                 `json:"timestamp"`
	Caller    string                 `json:"caller,omitempty"`
	Context   map[string]interface{} `json:"context,omitempty"`
}

// Logger wraps the standard log package with structured JSON logging.
type Logger struct {
	logger *log.Logger
}

// NewLogger initializes a new JSON logger.
func NewLogger() *Logger {
	return &Logger{logger: log.New(os.Stdout, "", 0)}
}

// logMessage creates a structured log entry.
func (l *Logger) logMessage(level, message string, ctx map[string]interface{}) {
	entry := LogEntry{
		Severity:  level,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
		Context:   ctx,
	}

	// Capture caller info
	if _, file, line, ok := runtime.Caller(2); ok {
		entry.Caller = fmt.Sprintf("%s:%d", file, line)
	}

	logData, _ := json.Marshal(entry)
	l.logger.Println(string(logData))
}

// Predefined logging methods with optional context
func (l *Logger) Default(msg string, ctx ...map[string]interface{}) {
	l.logMessage(DEFAULT, msg, getContext(ctx))
}
func (l *Logger) Debug(msg string, ctx ...map[string]interface{}) {
	l.logMessage(DEBUG, msg, getContext(ctx))
}
func (l *Logger) Info(msg string, ctx ...map[string]interface{}) {
	l.logMessage(INFO, msg, getContext(ctx))
}
func (l *Logger) Notice(msg string, ctx ...map[string]interface{}) {
	l.logMessage(NOTICE, msg, getContext(ctx))
}
func (l *Logger) Warning(msg string, ctx ...map[string]interface{}) {
	l.logMessage(WARNING, msg, getContext(ctx))
}
func (l *Logger) Error(msg string, ctx ...map[string]interface{}) {
	l.logMessage(ERROR, msg, getContext(ctx))
}
func (l *Logger) Critical(msg string, ctx ...map[string]interface{}) {
	l.logMessage(CRITICAL, msg, getContext(ctx))
}
func (l *Logger) Alert(msg string, ctx ...map[string]interface{}) {
	l.logMessage(ALERT, msg, getContext(ctx))
}
func (l *Logger) Emergency(msg string, ctx ...map[string]interface{}) {
	l.logMessage(EMERGENCY, msg, getContext(ctx))
}

// getContext extracts optional context from variadic arguments
func getContext(ctx []map[string]interface{}) map[string]interface{} {
	if len(ctx) > 0 {
		return ctx[0]
	}
	return nil
}
