package log

import (
	"fmt"
	"io"
	"os"
	"sync/atomic"
)

type Level uint32

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

type Logger struct {
	Out   io.Writer
	Level Level
}

func New() *Logger {
	return &Logger{
		Out:   os.Stderr,
		Level: DebugLevel,
	}
}

func (logger *Logger) Msgf(format string, args ...interface{}) {
	io.WriteString(os.Stdout, fmt.Sprintf(format, args...))
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	if logger.Level >= DebugLevel {
		var msg = "DEBUG ▶ "
		msg += fmt.Sprintf(format, args...)
		msg += "\n"
		io.WriteString(logger.Out, msg)
	}
}

func (logger *Logger) SetLevel(level Level) {
	atomic.StoreUint32((*uint32)(&logger.Level), uint32(level))
}
