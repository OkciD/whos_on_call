package sqldblogger_adapter

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/pkg/logger"

	sqldblogger "github.com/simukti/sqldb-logger"
)

type loggerAdapter struct {
	logger logger.Logger
}

func New(logger logger.Logger) sqldblogger.Logger {
	return &loggerAdapter{logger: logger}
}

func (l *loggerAdapter) Log(ctx context.Context, level sqldblogger.Level, msg string, data map[string]any) {
	logger := l.logger.WithContext(ctx).WithFields(data)

	switch level {
	case sqldblogger.LevelError:
		logger.Error(msg)
	case sqldblogger.LevelInfo:
		logger.Info(msg)
	case sqldblogger.LevelDebug:
		logger.Debug(msg)
	case sqldblogger.LevelTrace:
		logger.Trace(msg)
	default:
		logger.Debug(msg)
	}
}
