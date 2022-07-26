package logutil

import "time"

type Logger interface {
	Log(level LogLevel, body string) error
}

type Log struct {
	Level     LogLevel
	Timestamp time.Time
	Body      string
	Callers   []string
}

type LogLevel uint8

const (
	LogLevelUnknown = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelErr
	LogLevelPanic
)
