package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log the information
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns a logger
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

// Debugf formats and prints the message when log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {

	// Make sure if no output is provided then default stdout
	if l.output == nil {
		l.output = os.Stdout
	}

	// Implement me
	if l.threshold > LevelDebug {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

// Infof format and prints the message when log level is info or higher
func (l *Logger) Infof(format string, args ...any) {

	// Make sure if no output is provided then default stdout
	if l.output == nil {
		l.output = os.Stdout
	}

	// Implement me
	if l.threshold > LevelInfo {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

// Errorf format and prints the message when log level is error or higher
func (l *Logger) Errorf(format string, args ...any) {

	// Make sure if no output is provided then default stdout
	if l.output == nil {
		l.output = os.Stdout
	}

	// Implement me
	if l.threshold > LevelError {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}
