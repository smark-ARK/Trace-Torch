package pocketlog

// Level represents an available logging level.
type Level byte

const (
	// LevelDebug represents the lowest level of log. Mostly used for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information that is deemed important.
	LevelInfo
	// LevelWarn represents a warning. it needs attention but is not breaking anything
	LevelWarn
	// LevelError represents one of the most important levels of logging. Only to be used to trace errors.
	LevelError
	// LevelFatal represents the highest priority of logging. This means the server is broken and needs immediate attention.
	LevelFatal
)
