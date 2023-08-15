package pocketlog

// Level to represent error logs logging level
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents the logging level
	LevelInfo
	// LevelError represents the error level
	LevelError
)
