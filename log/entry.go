package log

import (
	"fmt"
)

type Entry struct {
	log  Logger
	opts options
}

var defaultEntry *Entry = New(&nopLogger{})

func New(log Logger, opts ...Option) *Entry {
	options := options{
		level: LevelDebug,
	}
	for _, opt := range opts {
		opt(&options)
	}
	return &Entry{
		opts: options,
		log:  log,
	}
}

func AsDefault(entry *Entry) {
	defaultEntry = entry
}

func GetDefault() *Entry {
	return defaultEntry
}

func With(kvpair ...interface{}) *Entry {
	return &Entry{
		log:  kv(defaultEntry.log, kvpair...),
		opts: defaultEntry.opts,
	}
}

func WithError(err error) *Entry {
	return &Entry{
		log:  kv(defaultEntry.log, errorKey, err),
		opts: defaultEntry.opts,
	}
}

func (entry *Entry) With(kvpair ...interface{}) *Entry {
	return &Entry{
		log:  kv(entry.log, kvpair...),
		opts: entry.opts,
	}
}

func (entry *Entry) WithError(err error) *Entry {
	return &Entry{
		log:  kv(entry.log, errorKey, err),
		opts: entry.opts,
	}
}

func (entry *Entry) Debug(args ...interface{}) {
	if entry.opts.level.Enabled(LevelDebug) {
		entry.log.Print(LevelDebug, 1, messageKey, fmt.Sprint(args...))
	}
}

func (entry *Entry) Debugf(format string, args ...interface{}) {
	if entry.opts.level.Enabled(LevelDebug) {
		entry.log.Print(LevelDebug, 1, messageKey, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Info(args ...interface{}) {
	if entry.opts.level.Enabled(LevelInfo) {
		entry.log.Print(LevelInfo, 1, messageKey, fmt.Sprint(args...))
	}
}

func (entry *Entry) Infof(format string, args ...interface{}) {
	if entry.opts.level.Enabled(LevelInfo) {
		entry.log.Print(LevelInfo, 1, messageKey, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Warn(args ...interface{}) {
	if entry.opts.level.Enabled(LevelWarn) {
		entry.log.Print(LevelWarn, 1, messageKey, fmt.Sprint(args...))
	}
}

func (entry *Entry) Warnf(format string, args ...interface{}) {
	if entry.opts.level.Enabled(LevelWarn) {
		entry.log.Print(LevelWarn, 1, messageKey, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Error(args ...interface{}) {
	if entry.opts.level.Enabled(LevelError) {
		entry.log.Print(LevelError, 1, messageKey, fmt.Sprint(args...))
	}
}

func (entry *Entry) Errorf(format string, args ...interface{}) {
	if entry.opts.level.Enabled(LevelError) {
		entry.log.Print(LevelError, 1, messageKey, fmt.Sprintf(format, args...))
	}
}
