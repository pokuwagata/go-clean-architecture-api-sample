package logger

import (
	"log"
	"io"
)

// NewLogger new log manager
func NewLogger(w io.Writer) *Logger {
	log.SetOutput(w)
	return &Logger{}
}

const (
	LogPrefixError = "[Error]"
	LogPrefixInfo = "[Info]"
	LogPrefixDebug = "[Debug]"
)

type Logger struct{}

func (l *Logger) Errorf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixError)
	log.Printf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixInfo)
	log.Printf(format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixDebug)
	log.Printf(format, args...)
}
