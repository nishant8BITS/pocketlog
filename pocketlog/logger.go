package pocketlog

// Logger is used to log the information
type Logger struct {
	threshold Level
}

// New returns a logger
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

// Debugf formats and prints the message when log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	// Implement me
}

// Infof format and prints the message when log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
	// Implement me
}

// Errorf format and prints the message when log level is error or higher
func (l *Logger) Errorf(format string, args ...any) {
	// Implement me
}
