package log

type Option func(*options)

type options struct {
	level Level
}

func AllowLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}
