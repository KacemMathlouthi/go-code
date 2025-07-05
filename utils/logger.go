package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// LogLevel represents the logging level
type LogLevel string

const (
	DEBUG   LogLevel = "DEBUG"
	INFO    LogLevel = "INFO"
	WARNING LogLevel = "WARNING"
	ERROR   LogLevel = "ERROR"
)

// LogEntry represents a structured log entry
type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     LogLevel               `json:"level"`
	Message   string                 `json:"message"`
	Category  string                 `json:"category,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

// Logger handles logging operations
type Logger struct {
	fileLogger *log.Logger
	file       *os.File
}

var (
	globalLogger *Logger
	logFile      = "app.log"
)

// InitLogger initializes the global logger
func InitLogger() error {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	globalLogger = &Logger{
		fileLogger: log.New(file, "", 0),
		file:       file,
	}

	// Log initialization
	globalLogger.Log(INFO, "Logger initialized", "system", nil)
	return nil
}

// CloseLogger closes the log file
func CloseLogger() {
	if globalLogger != nil && globalLogger.file != nil {
		globalLogger.file.Close()
	}
}

// Log writes a structured log entry
func (l *Logger) Log(level LogLevel, message, category string, data map[string]interface{}) {
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Message:   message,
		Category:  category,
		Data:      data,
	}

	// Convert to JSON
	jsonData, err := json.Marshal(entry)
	if err != nil {
		// Fallback to simple logging if JSON marshaling fails
		l.fileLogger.Printf("[%s] %s: %s", level, category, message)
		return
	}

	// Write to file
	l.fileLogger.Println(string(jsonData))

	// Also print to console for development
	fmt.Printf("[%s] %s: %s\n", level, category, message)
	if data != nil {
		fmt.Printf("  Data: %+v\n", data)
	}
}

// Convenience methods for different log levels
func (l *Logger) Debug(message, category string, data map[string]interface{}) {
	l.Log(DEBUG, message, category, data)
}

func (l *Logger) Info(message, category string, data map[string]interface{}) {
	l.Log(INFO, message, category, data)
}

func (l *Logger) Warning(message, category string, data map[string]interface{}) {
	l.Log(WARNING, message, category, data)
}

func (l *Logger) Error(message, category string, data map[string]interface{}) {
	l.Log(ERROR, message, category, data)
}

// Global convenience functions
func LogDebug(message, category string, data map[string]interface{}) {
	if globalLogger != nil {
		globalLogger.Debug(message, category, data)
	}
}

func LogInfo(message, category string, data map[string]interface{}) {
	if globalLogger != nil {
		globalLogger.Info(message, category, data)
	}
}

func LogWarning(message, category string, data map[string]interface{}) {
	if globalLogger != nil {
		globalLogger.Warning(message, category, data)
	}
}

func LogError(message, category string, data map[string]interface{}) {
	if globalLogger != nil {
		globalLogger.Error(message, category, data)
	}
}

// Specific logging functions for LLM and tools
func LogLLMRequest(prompt string, model string) {
	LogInfo("LLM Request", "llm", map[string]interface{}{
		"prompt": prompt,
		"model":  model,
	})
}

func LogLLMResponse(response string, model string, duration time.Duration) {
	LogInfo("LLM Response", "llm", map[string]interface{}{
		"response":    response,
		"model":       model,
		"duration":    duration.String(),
		"duration_ms": duration.Milliseconds(),
	})
}

func LogToolCall(toolName string, toolArgs map[string]string) {
	LogInfo("Tool Call", "tool", map[string]interface{}{
		"tool_name": toolName,
		"tool_args": toolArgs,
	})
}

func LogToolResult(toolName string, result string, err error) {
	data := map[string]interface{}{
		"tool_name": toolName,
		"result":    result,
	}

	if err != nil {
		data["error"] = err.Error()
		LogError("Tool Execution Failed", "tool", data)
	} else {
		LogInfo("Tool Execution Success", "tool", data)
	}
}
