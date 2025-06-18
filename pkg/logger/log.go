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

var Log *logrus.Logger

func Init(level logrus.Level) {
	Log = logrus.New()

	Log.SetLevel(level)

	Log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

func Fatal(textMessage string, user string, additionalTextMessage string, location string, err error) {
	Log.WithFields(logrus.Fields{
		"user":     user,
		"text":     additionalTextMessage,
		"location": location,
		"err":      err,
	}).Fatal()
}

func Error(textMessage string, user string, additionalTextMessage string, location string, err error) {
	Log.WithFields(logrus.Fields{
		"user":     user,
		"text":     additionalTextMessage,
		"location": location,
		"err":      err,
	}).Error()
}

func Warn(textMessage string, user string, additionalTextMessage string, location string, err error) {
	Log.WithFields(logrus.Fields{
		"user":     user,
		"text":     additionalTextMessage,
		"location": location,
		"err":      err,
	}).Warn()
}

func Info(textMessage string, user string, additionalTextMessage string, location string, err error) {
	Log.WithFields(logrus.Fields{
		"user":     user,
		"text":     additionalTextMessage,
		"location": location,
		"err":      err,
	}).Info()
}

func Debug(textMessage string, user string, additionalTextMessage string, location string, err error) {
	Log.WithFields(logrus.Fields{
		"user":     user,
		"text":     additionalTextMessage,
		"location": location,
		"err":      err,
	}).Debug()
}
