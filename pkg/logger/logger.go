package logger

import (
	"context"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/appcontext"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/constant"
	"log"
	"os"

	"git.thoughtworks.net/mahadeva/sample-golang/pkg/config"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

type ContextLogger struct {
	entry *logrus.Entry
}

var NonContext *ContextLogger

func Setup() {
	level, err := logrus.ParseLevel(config.Log().Level())
	if err != nil {
		log.Fatalf(err.Error())
	}

	logger = &logrus.Logger{
		Out:   os.Stdout,
		Hooks: make(logrus.LevelHooks),
		Level: level,
	}

	if config.Log().Format() == "text" {
		logger.Formatter = &logrus.TextFormatter{}
	} else {
		logger.Formatter = &logrus.JSONFormatter{}
	}
	NonContext = &ContextLogger{
		entry: logrus.NewEntry(logger),
	}
}

func NewContextLogger(ctx context.Context) *ContextLogger {
	logFields := logger.WithFields(logrus.Fields{
		constant.CorrelationIDHeader: appcontext.GetCorrelationID(ctx),
	})

	return &ContextLogger{
		entry: logFields,
	}
}

func (l *ContextLogger) Error(method, errMessage string, err error) {
	l.entry.
		WithField(constant.ContextLogTag, method).
		WithField(constant.ErrorLogTag, err).
		Error(errMessage)
}

func (l *ContextLogger) Info(method, msg string, fields map[string]interface{}) {
	if fields != nil {
		for key, val := range fields {
			l.entry = l.entry.WithField(key, val)
		}
	}
	l.entry.WithField(constant.ContextLogTag, method).Info(msg)
}

func (l *ContextLogger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *ContextLogger) WarnWithFields(fields map[string]interface{}, args ...interface{}) {
	if fields != nil {
		for key, val := range fields {
			l.entry = l.entry.WithField(key, val)
		}
	}
	l.entry.Warn(args...)
}
