package pocketlog

// Logger is used to log information.
type Logger struct {
	threshold Level
}

func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

func (l *Logger) Debugf(format string, args ...any) {
	// Yet to be implemented
}
func (l *Logger) Infof(format string, args ...any) {
	// Yet to be implemented
}
func (l *Logger) Warnf(format string, args ...any) {
	// Yet to be implemented
}
func (l *Logger) Errorf(format string, args ...any) {
	// Yet to be implemented
}
func (l *Logger) Fatalf(format string, args ...any) {
	// Yet to be implemented
}
