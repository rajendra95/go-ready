package logger

import (
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

type Fields map[string]interface{}

var logger *logrus.Logger

func InitLogger(logLevel string) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.WarnLevel
	}

	logger = &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			FullTimestamp:          true,
			DisableLevelTruncation: true,
		},
		Hooks: make(logrus.LevelHooks),
		Level: level,
	}
}

func WithFields(fields Fields) Logger {
	if logger == nil {
		panic("Logger not initialized")
	}

	srcInfo := getSourceInfoFields()
	if fields != nil {
		for key, val := range fields {
			srcInfo[key] = val
		}
	}
	return logger.WithFields(logrus.Fields(srcInfo))
}

func getSourceInfoFields() map[string]interface{} {
	file, line := getFileInfo(4)
	m := map[string]interface{}{
		"file": file,
		"line": line,
	}
	return m
}

func getFileInfo(subtractStackLevels int) (string, int) {
	_, file, line, _ := runtime.Caller(subtractStackLevels)
	return chopPath(file), line
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i != -1 {
		return original[i+1:]
	}
	return original
}

func Debugf(format string, args ...interface{}) {
	WithFields(nil).Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	WithFields(nil).Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	WithFields(nil).Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	WithFields(nil).Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	WithFields(nil).Fatalf(format, args...)
}

func Debug(args ...interface{}) {
	WithFields(nil).Debug(args...)
}

func Info(args ...interface{}) {
	WithFields(nil).Info(args...)
}

func Warn(args ...interface{}) {
	WithFields(nil).Warn(args...)
}

func Error(args ...interface{}) {
	WithFields(nil).Error(args...)
}

func Fatal(args ...interface{}) {
	WithFields(nil).Fatal(args...)
}
