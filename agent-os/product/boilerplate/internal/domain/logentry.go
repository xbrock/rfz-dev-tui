// Package domain contains the core domain entities and value objects.
//
// This file defines the LogEntry entity which represents a single
// line of build output from Maven or other external processes.
package domain

import "time"

// LogLevel represents the severity level of a log entry.
type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

// String returns the string representation of a log level.
func (l LogLevel) String() string {
	switch l {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarning:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// LogEntry represents a single log line from a build process.
//
// Log entries are entities because they have a unique identity (ID)
// and their state can change over time (e.g., being marked as read).
type LogEntry struct {
	// ID is the unique identifier for this log entry.
	ID string

	// Timestamp is when the log was generated.
	Timestamp time.Time

	// Level is the severity level of the log.
	Level LogLevel

	// Message is the log message content.
	Message string

	// ComponentID is the ID of the component that generated this log.
	// Empty if not associated with a specific component.
	ComponentID string

	// BuildID is the ID of the build that generated this log.
	BuildID string

	// Source indicates the source of the log (maven, git, etc.)
	Source string
}

// NewLogEntry creates a new LogEntry with the current timestamp.
func NewLogEntry(level LogLevel, message string) LogEntry {
	return LogEntry{
		ID:        generateID(), // TODO: Implement ID generation
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
	}
}

// WithComponent returns a new LogEntry with the component ID set.
func (e LogEntry) WithComponent(componentID string) LogEntry {
	e.ComponentID = componentID
	return e
}

// WithBuild returns a new LogEntry with the build ID set.
func (e LogEntry) WithBuild(buildID string) LogEntry {
	e.BuildID = buildID
	return e
}

// IsError returns true if this is an error-level log.
func (e LogEntry) IsError() bool {
	return e.Level == LogLevelError
}

// IsWarning returns true if this is a warning-level log.
func (e LogEntry) IsWarning() bool {
	return e.Level == LogLevelWarning
}

// FormatTimestamp returns the formatted timestamp.
func (e LogEntry) FormatTimestamp() string {
	return e.Timestamp.Format("15:04:05.000")
}

// generateID generates a unique ID for a log entry.
// TODO: Implement proper ID generation (UUID or sequential).
func generateID() string {
	return time.Now().Format("20060102150405.000")
}

// BuildResult represents the result of a build operation.
type BuildResult struct {
	// BuildID is the unique identifier for this build.
	BuildID string

	// ComponentID is the ID of the component that was built.
	ComponentID string

	// Success indicates whether the build was successful.
	Success bool

	// StartTime is when the build started.
	StartTime time.Time

	// EndTime is when the build completed.
	EndTime time.Time

	// ExitCode is the Maven process exit code.
	ExitCode int

	// ErrorMessage contains error details if the build failed.
	ErrorMessage string

	// Logs contains the build log entries.
	Logs []LogEntry
}

// Duration returns the build duration.
func (r BuildResult) Duration() time.Duration {
	return r.EndTime.Sub(r.StartTime)
}
