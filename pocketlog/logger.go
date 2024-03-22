package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns you a logger, ready to log at the required threshold.
// The default output is Stdout.
func New(threshold Level, opts ...Option) *Logger {
	l := &Logger{
		threshold: threshold,
		output:    os.Stdout,
	}

	for _, configFunc := range opts {
		configFunc(l)
	}
	return l

}

// Prints the message to the output
func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

// * Debug formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {

	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelDebug {
		return
	}

	l.logf(format, args...)
}
func (l *Logger) Infof(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelInfo {
		return
	}
	l.logf(format, args...)
}
func (l *Logger) Warnf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelWarn {
		return
	}
	l.logf(format, args...)
}
func (l *Logger) Errorf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelError {
		return
	}
	l.logf(format, args...)
}
func (l *Logger) Fatalf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelFatal {
		return
	}
	l.logf(format, args...)
}
