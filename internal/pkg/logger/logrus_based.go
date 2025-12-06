package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type logrusBasedLogger struct {
	entry *logrus.Entry
}

func (l logrusBasedLogger) WithField(key string, value any) Logger {
	return logrusBasedLogger{entry: l.entry.WithField(key, value)}
}

func (l logrusBasedLogger) WithFields(fields Fields) Logger {
	return logrusBasedLogger{entry: l.entry.WithFields(logrus.Fields(fields))}
}

func (l logrusBasedLogger) WithError(err error) Logger {
	return logrusBasedLogger{entry: l.entry.WithError(err)}
}

func (l logrusBasedLogger) Trace(msg string) {
	l.entry.Trace(msg)
}

func (l logrusBasedLogger) Debug(msg string) {
	l.entry.Debug(msg)
}

func (l logrusBasedLogger) Info(msg string) {
	l.entry.Info(msg)
}

func (l logrusBasedLogger) Warn(msg string) {
	l.entry.Warn(msg)
}

func (l logrusBasedLogger) Error(msg string) {
	l.entry.Error(msg)
}

func (l logrusBasedLogger) Fatal(msg string) {
	l.entry.Fatal(msg)
}

func (l logrusBasedLogger) Panic(msg string) {
	l.entry.Panic(msg)
}

func NewLogrusBasedLogger(cfg *Config) Logger {
	logrusLogger := logrus.New()

	if cfg.Format == LogFormatText {
		logrusLogger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		})
	} else {
		logrusLogger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	}

	logLevel, err := logrus.ParseLevel(string(cfg.Level))
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logrusLogger.SetLevel(logLevel)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unk"
	}
	logrusEntry := logrusLogger.WithField("host", hostname)

	return &logrusBasedLogger{
		entry: logrusEntry,
	}
}
