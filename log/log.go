package log

const (
	messageKey = "message"
	errorKey   = "error"
)

// Logger logger interface
type Logger interface {
	Print(level Level, depth int, kvpair ...interface{})
}

// kvLogger nested logger with kvpair
type kvLogger struct {
	// log interface
	log Logger
	//kvpair always be logged
	kvpair []interface{}
}

// newKvLogger new nested kvlogger
func newKvLogger(log Logger, kvpair ...interface{}) *kvLogger {
	return &kvLogger{log: log, kvpair: kvpair}
}

func (l *kvLogger) Print(level Level, depth int, kvpair ...interface{}) {
	l.log.Print(level, depth+1, append(l.kvpair, kvpair...)...)
}

func kv(log Logger, kvpair ...interface{}) Logger {
	return newKvLogger(log, kvpair...)
}

func Debug(args ...interface{}) {
	defaultEntry.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	defaultEntry.Debugf(format, args...)
}

func Info(args ...interface{}) {
	defaultEntry.Info(args...)
}

func Infof(format string, args ...interface{}) {
	defaultEntry.Infof(format, args...)
}

func Warn(args ...interface{}) {
	defaultEntry.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	defaultEntry.Warnf(format, args...)
}

func Error(args ...interface{}) {
	defaultEntry.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	defaultEntry.Errorf(format, args...)
}
