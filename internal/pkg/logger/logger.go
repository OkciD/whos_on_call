package logger

import (
	"context"
	"maps"
)

type Fields map[string]any

func (f Fields) Extend(fields Fields) {
	maps.Copy(f, fields)
}

type Logger interface {
	WithField(key string, value any) Logger
	WithFields(fields Fields) Logger
	WithError(err error) Logger
	WithContext(ctx context.Context) Logger

	Trace(msg string)
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	Panic(msg string)
}
