package logger

import (
	"github.com/sirupsen/logrus"
)

// PanicLevel 0
// FatalLevel 1
// ErrorLevel 2
// WarnLevel 3
// InfoLevel 4
// DebugLevel 5
// TraceLevel 6

type Logger struct {
	logger *logrus.Logger
}

func NewLogger(level logrus.Level) Logger {
	logger := logrus.New()

	logger.SetLevel(level)

	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	return Logger{
		logger: logger,
	}
}

func (l *Logger) Fatal(textMessage string, user string, additionalTextMessage string, location string, err error) {
	l.logger.WithFields(logrus.Fields{
		"user":     user,
		"text":     additionalTextMessage,
		"location": location,
		"err":      err,
	}).Fatal()
}

func (l *Logger) Error(textMessage string, user string, additionalTextMessage string, location string, err error) {
	l.logger.WithFields(logrus.Fields{
		"user":     user,
		"text":     additionalTextMessage,
		"location": location,
		"err":      err,
	}).Error()
}

func (l *Logger) Warn(textMessage string, user string, additionalTextMessage string, location string, err error) {
	l.logger.WithFields(logrus.Fields{
		"user":     user,
		"text":     additionalTextMessage,
		"location": location,
		"err":      err,
	}).Warn()
}

func (l *Logger) Info(textMessage string, user string, additionalTextMessage string, location string, err error) {
	l.logger.WithFields(logrus.Fields{
		"user":     user,
		"text":     additionalTextMessage,
		"location": location,
		"err":      err,
	}).Info()
}

func (l *Logger) Debug(textMessage string, user string, additionalTextMessage string, location string, err error) {
	l.logger.WithFields(logrus.Fields{
		"user":     user,
		"text":     additionalTextMessage,
		"location": location,
		"err":      err,
	}).Debug()
}
