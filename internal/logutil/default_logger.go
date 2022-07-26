package logutil

import (
	"encoding/json"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type DefaultLogger struct {
	config DefaultLoggerConfig
}

type DefaultLoggerConfig struct {
	InfoWriter io.Writer
	ErrWriter  io.Writer
	MaxCallers int
}

// NewDefaultLogger returns a new DefaultLogger.
// It also sets default writers if needed
func NewDefaultLogger(config DefaultLoggerConfig) *DefaultLogger {
	if config.ErrWriter == nil {
		config.ErrWriter = os.Stderr
	}
	if config.InfoWriter == nil {
		config.InfoWriter = os.Stdout
	}
	if config.MaxCallers <= 0 {
		config.MaxCallers = 1
	}

	return &DefaultLogger{
		config: config,
	}
}

func (s *DefaultLogger) Log(level LogLevel, body string) error {
	callers := []string{}
	for i := 1; i < s.config.MaxCallers+1; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		file = strings.TrimPrefix(file, "/usr/local/go/app/")
		caller := runtime.FuncForPC(pc).Name() + " (at " + file + ":" + strconv.Itoa(line) + ")"
		callers = append(callers, caller)
	}

	l := &Log{
		Timestamp: time.Now(),
		Callers:   callers,
		Level:     level,
		Body:      body,
	}

	if l.Level >= LogLevelErr {
		return json.NewEncoder(s.config.ErrWriter).Encode(l) // write errors and panics to stderr writer
	}
	return json.NewEncoder(s.config.InfoWriter).Encode(l) // write info and warning to stdout writer
}
