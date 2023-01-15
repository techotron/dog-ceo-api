package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetOutput(os.Stdout)
}

// Info wrapper of logrus
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// Infof wrapper of logrus
func Infof(message string, args ...interface{}) {
	Logger.Infof(message, args...)
}

// Debug wrapper of logrus
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

// Debugf wrapper of logrus
func Debugf(message string, args ...interface{}) {
	Logger.Debugf(message, args...)
}

// Warn wrapper of logrus
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// Warnf wrapper of logrus
func Warnf(message string, args ...interface{}) {
	Logger.Warnf(message, args...)
}

// Fatal wrapper of logrus
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// Fatalf wrapper of logrus
func Fatalf(message string, args ...interface{}) {
	Logger.Fatalf(message, args...)
}

// Error wrapper of logrus
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// Errorf wrapper of logrus
func Errorf(message string, args ...interface{}) {
	Logger.Errorf(message, args...)
}
