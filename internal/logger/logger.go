package logger

import (
	"fmt"
	"os"
)

// Logger provides structured logging with dry-run support
type Logger struct {
	DryRun  bool
	Verbose bool
}

// New creates a new logger instance
func New(dryRun, verbose bool) *Logger {
	return &Logger{
		DryRun:  dryRun,
		Verbose: verbose,
	}
}

// Info logs an informational message
func (l *Logger) Info(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format+"\n", args...)
}

// Success logs a success message
func (l *Logger) Success(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, "✓ "+format+"\n", args...)
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "✗ "+format+"\n", args...)
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, "⚠ "+format+"\n", args...)
}

// Debug logs a debug message (only if verbose is enabled)
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.Verbose {
		fmt.Fprintf(os.Stdout, "[DEBUG] "+format+"\n", args...)
	}
}

// DryRunMsg logs a message indicating what would be done in dry-run mode
func (l *Logger) DryRunMsg(format string, args ...interface{}) {
	if l.DryRun {
		fmt.Fprintf(os.Stdout, "[DRY-RUN] "+format+"\n", args...)
	}
}

// Action logs an action that will be taken
func (l *Logger) Action(format string, args ...interface{}) {
	if l.DryRun {
		fmt.Fprintf(os.Stdout, "[DRY-RUN] Would: "+format+"\n", args...)
	} else {
		fmt.Fprintf(os.Stdout, "→ "+format+"\n", args...)
	}
}

// Step logs a step in a multi-step process
func (l *Logger) Step(step int, total int, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stdout, "[%d/%d] %s\n", step, total, msg)
}
