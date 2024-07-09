package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
)

const (
	ErrorLevel uint32 = iota
	InfoLevel
	DebugLevel
)

var (
	Default = New(DebugLevel)
	mu      sync.Mutex
)

type logger struct {
	prefix string
	level  uint32
	w      io.Writer
}

func New(level uint32) *logger {
	return &logger{
		level: level,
		w:     os.Stderr,
	}
}

func NewWithPrefix(level uint32, prefix string) *logger {
	return &logger{
		prefix: prefix,
		level:  level,
		w:      os.Stderr,
	}
}

func (l *logger) Error(format string, args ...any) {
	l.log(ErrorLevel, format, args...)
}

func (l *logger) Info(format string, args ...any) {
	l.log(InfoLevel, format, args...)
}

func (l *logger) Debug(format string, args ...any) {
	l.log(DebugLevel, format, args...)
}

func (l *logger) isLevelEnabled(level uint32) bool {
	return level <= l.level
}

func (l *logger) log(level uint32, format string, args ...any) {
	if l == nil {
		return
	}

	if !l.isLevelEnabled(level) {
		return
	}
	mu.Lock()
	defer mu.Unlock()

	l.w.Write([]byte(fmt.Sprintf("%s %s\n", l.prefix, fmt.Sprintf(format, args...))))
}
