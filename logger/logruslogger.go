package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type logruslogger struct {
	log *logrus.Logger
}

func NewLogrusLogger() Logger {
	log := logrus.New()
	log.Out = os.Stdout
	return &logruslogger{log: log}
}

func (l *logruslogger) Trace(args ...any) {
	l.log.Trace(args...)
}
func (l *logruslogger) Tracef(format string, args ...any) {
	l.log.Tracef(format, args...)
}
func (l *logruslogger) Debug(args ...any) {
	l.log.Debug(args...)
}
func (l *logruslogger) Debugf(format string, args ...any) {
	l.log.Debugf(format, args...)
}
func (l *logruslogger) Info(args ...any) {
	l.log.Info(args...)
}
func (l *logruslogger) Infof(format string, args ...any) {
	l.log.Infof(format, args...)
}
func (l *logruslogger) Warn(args ...any) {
	l.log.Warn(args...)
}
func (l *logruslogger) Warnf(format string, args ...any) {
	l.log.Warnf(format, args...)
}
func (l *logruslogger) Error(args ...any) {
	l.log.Error(args...)
}
func (l *logruslogger) Errorf(format string, args ...any) {
	l.log.Errorf(format, args...)
}
func (l *logruslogger) Fatal(args ...any) {
	l.log.Fatal(args...)
}
func (l *logruslogger) Fatalf(format string, args ...any) {
	l.log.Fatalf(format, args...)
}
